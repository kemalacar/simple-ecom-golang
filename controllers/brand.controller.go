package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kemalacar/go-ecom/models"
	"github.com/kemalacar/go-ecom/s3"
	"github.com/kemalacar/go-ecom/utils"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type BrandController struct {
	DB *gorm.DB
}

func NewBrandController(DB *gorm.DB) BrandController {
	return BrandController{DB}
}

func (bc *BrandController) GetAll(ctx *gin.Context) {

	intLimit, offset := utils.GetLimitOffset(ctx)

	var brands []models.Brand
	results := bc.DB.Limit(intLimit).Offset(offset).Find(&brands)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(brands), "data": models.BrandsToDto(brands)})

}

func (bc *BrandController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	result := bc.DB.Delete(&models.Brand{}, "id = ?", id)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No brand with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "result": true})
}

func (bc *BrandController) Create(ctx *gin.Context) {
	//currentUser := ctx.MustGet("currentUser").(models.User)
	var payload models.BrandDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newBrand := models.BrandToModel(&payload)

	result := bc.DB.Create(&newBrand)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Brand with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": models.BrandToDto(newBrand)})
}

func (bc *BrandController) Update(ctx *gin.Context) {

	//id := ctx.Param("id")
	//currentUser := ctx.MustGet("currentUser").(models.User)

	ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Not implemented yet"})
}

func (bc *BrandController) Test(ctx *gin.Context) {

	dto := models.BrandDto{
		Id:      0,
		LogoUrl: "",
		Name:    "",
	}
	s3.MyTest(&dto)
	ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": dto})
}
