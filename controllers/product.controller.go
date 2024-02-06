package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kemalacar/go-ecom/initializers"
	"github.com/kemalacar/go-ecom/models"
	"github.com/kemalacar/go-ecom/s3"
	"github.com/kemalacar/go-ecom/utils"
	format_errors "github.com/kemalacar/go-ecom/utils/format-errors"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type ProductController struct {
	DB *gorm.DB
}

func NewProductController(DB *gorm.DB) ProductController {
	return ProductController{DB}
}

func (pc *ProductController) GetAll(ctx *gin.Context) {

	intLimit, offset := utils.GetLimitOffset(ctx)
	var Products []models.Product
	results := pc.DB.Limit(intLimit).Offset(offset).Preload("Brand").
		Preload("StoreProducts").Preload("Images").Preload("StoreProducts.Images").
		Preload("StoreProducts.SizeOption").Preload("StoreProducts.ColorOption").Find(&Products)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(Products), "data": models.ProductsToDto(Products)})

}

func (pc *ProductController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var product models.Product
	// Find the post
	if err := initializers.DB.Unscoped().Preload("StoreProducts").Preload("Images").Preload("StoreProducts.Images").First(&product, id).Error; err != nil {
		format_errors.RecordNotFound(ctx, err)
		return
	}

	var imageList []string

	for _, image := range product.Images {
		imageList = append(imageList, image.Big[strings.LastIndex(image.Big, "/")+1:])
	}

	for _, sp := range product.StoreProducts {
		for _, image := range sp.Images {
			imageList = append(imageList, image.Big[strings.LastIndex(image.Big, "/")+1:])
		}
	}

	s3.DeleteObjects(ctx, imageList)

	//result := pc.DB.Select(clause.Associations).Delete(&product) // SOFT DELETE
	result := pc.DB.Select("StoreProducts").Unscoped().Delete(&product)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Product with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "result": true})
}

func (pc *ProductController) Create(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	var cpr models.CreateProductRequest
	if err := ctx.ShouldBindJSON(&cpr); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newProduct := models.ProductToModel(&cpr)

	for i := range newProduct.StoreProducts {
		storeProduct := newProduct.StoreProducts[i]
		storeProduct.StoreId = currentUser.StoreId
	}

	s3.UploadBase64(ctx, &newProduct)

	result := pc.DB.Create(&newProduct)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Product with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": models.ProductToDto(newProduct)})
}
func (pc *ProductController) Update(ctx *gin.Context) {
	//id := ctx.Param("id")
	//currentUser := ctx.MustGet("currentUser").(models.User)

	ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Not implemented yet"})
}

func (pc *ProductController) GetPreSignUrl(ctx *gin.Context) {

	url := s3.GetPreSignURL(ctx)

	ctx.JSON(http.StatusConflict, gin.H{"status": "success", "data": url})
}
