package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kemalacar/go-ecom/controllers"
	"github.com/kemalacar/go-ecom/initializers"
	"github.com/kemalacar/go-ecom/routes"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	PostController      controllers.PostController
	PostRouteController routes.PostRouteController

	BrandController      controllers.BrandController
	BrandRouteController routes.BrandRouteController

	StoreController      controllers.StoreController
	StoreRouteController routes.StoreRouteController

	ProductOptionController      controllers.ProductOptionController
	ProductOptionRouteController routes.ProductOptionRouteController

	ProductController      controllers.ProductController
	ProductRouteController routes.ProductRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	PostController = controllers.NewPostController(initializers.DB)
	PostRouteController = routes.NewRoutePostController(PostController)

	BrandController = controllers.NewBrandController(initializers.DB)
	BrandRouteController = routes.NewRouteBrandController(BrandController)

	StoreController = controllers.NewStoreController(initializers.DB)
	StoreRouteController = routes.NewRouteStoreController(StoreController)

	ProductOptionController = controllers.NewProductOptionController(initializers.DB)
	ProductOptionRouteController = routes.NewRouteProductOptionController(ProductOptionController)

	ProductController = controllers.NewProductController(initializers.DB)
	ProductRouteController = routes.NewRouteProductController(ProductController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	sess := ConnectAws(config)

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	router.Use(func(c *gin.Context) {
		c.Set("sess", sess)
		c.Next()
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	PostRouteController.PostRoute(router)
	BrandRouteController.BrandRoute(router)
	StoreRouteController.StoreRoute(router)
	ProductOptionRouteController.ProductOptionRoute(router)
	ProductRouteController.ProductRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}

var AccessKeyID string
var SecretAccessKey string
var MyRegion string

func ConnectAws(config initializers.Config) *session.Session {
	AccessKeyID = config.S3accessKey
	SecretAccessKey = config.S3secretKey
	MyRegion = "us-east-2"
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}
