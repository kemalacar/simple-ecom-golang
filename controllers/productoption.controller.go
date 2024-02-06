package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kemalacar/go-ecom/models"
	"github.com/kemalacar/go-ecom/utils"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type ProductOptionController struct {
	DB *gorm.DB
}

func NewProductOptionController(DB *gorm.DB) ProductOptionController {
	return ProductOptionController{DB}
}

func (poc *ProductOptionController) GetAll(ctx *gin.Context) {

	intLimit, offset := utils.GetLimitOffset(ctx)

	var ProductOptions []models.ProductOption
	results := poc.DB.Limit(intLimit).Offset(offset).Find(&ProductOptions)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(ProductOptions), "data": models.ProductOptionsToDto(ProductOptions)})

}

func (poc *ProductOptionController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	result := poc.DB.Delete(&models.ProductOption{}, "id = ?", id)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No ProductOption with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "result": true})
}

func (poc *ProductOptionController) Create(ctx *gin.Context) {
	//currentUser := ctx.MustGet("currentUser").(models.User)
	var payload models.CreateProductOptionRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newProductOption := models.ProductOptionToModel(&payload)

	result := poc.DB.Create(&newProductOption)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "ProductOption with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": models.ProductOptionToDto(newProductOption)})
}

func (poc *ProductOptionController) Update(ctx *gin.Context) {

	//id := ctx.Param("id")
	//currentUser := ctx.MustGet("currentUser").(models.User)

	ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Not implemented yet"})
}
