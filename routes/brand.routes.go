package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kemalacar/go-ecom/controllers"
	"github.com/kemalacar/go-ecom/middleware"
)

type BrandRouteController struct {
	brandController controllers.BrandController
}

func NewRouteBrandController(brandController controllers.BrandController) BrandRouteController {
	return BrandRouteController{brandController}
}

func (bc *BrandRouteController) BrandRoute(rg *gin.RouterGroup) {

	router := rg.Group("brand")
	router.Use(middleware.DeserializeUser())
	router.GET("/all", bc.brandController.GetAll)
	router.POST("/", bc.brandController.Create)
	router.DELETE("/:id", bc.brandController.Delete)
	router.PUT("/:id", bc.brandController.Update)
	router.GET("/test", bc.brandController.Test)
}
