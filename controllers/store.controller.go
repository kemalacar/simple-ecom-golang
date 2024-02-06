package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kemalacar/go-ecom/models"
	"github.com/kemalacar/go-ecom/utils"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type StoreController struct {
	DB *gorm.DB
}

func NewStoreController(DB *gorm.DB) StoreController {
	return StoreController{DB}
}

func (sc *StoreController) GetAll(ctx *gin.Context) {

	intLimit, offset := utils.GetLimitOffset(ctx)

	var stores []models.Store
	results := sc.DB.Limit(intLimit).Offset(offset).Find(&stores)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(stores), "data": models.StoresToDto(stores)})

}

func (sc *StoreController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	result := sc.DB.Delete(&models.Store{}, "id = ?", id)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Store with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "result": true})
}

func (sc *StoreController) Create(ctx *gin.Context) {
	//currentUser := ctx.MustGet("currentUser").(models.User)
	var payload models.CreateStoreRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newStore := models.StoreToModel(&payload)

	result := sc.DB.Create(&newStore)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Store with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": models.StoreToDto(newStore)})
}
func (sc *StoreController) Update(ctx *gin.Context) {
	//id := ctx.Param("id")
	//currentUser := ctx.MustGet("currentUser").(models.User)

	ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Not implemented yet"})
}
