package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kemalacar/go-ecom/controllers"
	"github.com/kemalacar/go-ecom/middleware"
)

type StoreRouteController struct {
	StoreController controllers.StoreController
}

func NewRouteStoreController(StoreController controllers.StoreController) StoreRouteController {
	return StoreRouteController{StoreController}
}

func (sc *StoreRouteController) StoreRoute(rg *gin.RouterGroup) {

	router := rg.Group("store")
	router.Use(middleware.DeserializeUser())
	router.GET("/all", sc.StoreController.GetAll)
	router.POST("/", sc.StoreController.Create)
	router.DELETE("/:id", sc.StoreController.Delete)
	router.PUT("/:id", sc.StoreController.Update)
}
