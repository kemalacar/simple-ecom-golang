package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kemalacar/go-ecom/controllers"
	"github.com/kemalacar/go-ecom/middleware"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("user")
	router.Use(middleware.DeserializeUser())
	router.GET("/me", uc.userController.GetMe)
	router.GET("/all", uc.userController.GetAll)
	router.DELETE("/:id", uc.userController.Delete)
	router.PUT("/:id", uc.userController.Update)
}
