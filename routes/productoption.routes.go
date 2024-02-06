package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kemalacar/go-ecom/controllers"
	"github.com/kemalacar/go-ecom/middleware"
)

type ProductOptionRouteController struct {
	ProductOptionController controllers.ProductOptionController
}

func NewRouteProductOptionController(ProductOptionController controllers.ProductOptionController) ProductOptionRouteController {
	return ProductOptionRouteController{ProductOptionController}
}

func (poc *ProductOptionRouteController) ProductOptionRoute(rg *gin.RouterGroup) {

	router := rg.Group("product-option")
	router.Use(middleware.DeserializeUser())
	router.GET("/all", poc.ProductOptionController.GetAll)
	router.POST("/", poc.ProductOptionController.Create)
	router.DELETE("/:id", poc.ProductOptionController.Delete)
	router.PUT("/:id", poc.ProductOptionController.Update)
}
