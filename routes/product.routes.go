package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kemalacar/go-ecom/controllers"
	"github.com/kemalacar/go-ecom/middleware"
)

type ProductRouteController struct {
	ProductController controllers.ProductController
}

func NewRouteProductController(ProductController controllers.ProductController) ProductRouteController {
	return ProductRouteController{ProductController}
}

func (prc *ProductRouteController) ProductRoute(rg *gin.RouterGroup) {

	router := rg.Group("product")
	router.Use(middleware.DeserializeUser())
	router.GET("/all", prc.ProductController.GetAll)
	router.POST("/", prc.ProductController.Create)
	router.DELETE("/:id", prc.ProductController.Delete)
	router.PUT("/:id", prc.ProductController.Update)
	router.GET("/get-pre-sign-url/:name", prc.ProductController.GetPreSignUrl)
}
