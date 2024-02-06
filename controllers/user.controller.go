package controllers

import (
	"github.com/kemalacar/go-ecom/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemalacar/go-ecom/models"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := models.UserToDto(currentUser)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

func (uc *UserController) GetAll(ctx *gin.Context) {

	intLimit, offset := utils.GetLimitOffset(ctx)

	var users []models.User
	results := uc.DB.Limit(intLimit).Offset(offset).Find(&users)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(users), "data": models.UsersToDto(users)})

}

func (uc *UserController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	result := uc.DB.Delete(&models.User{}, "id = ?", id)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No User with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "result": true})
}

func (uc *UserController) Update(ctx *gin.Context) {

	//id := ctx.Param("id")
	//currentUser := ctx.MustGet("currentUser").(models.User)

	ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Not implemented yet"})
}
