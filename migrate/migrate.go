package main

import (
	"fmt"
	"log"

	"github.com/kemalacar/go-ecom/initializers"
	"github.com/kemalacar/go-ecom/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Product{}, &models.Store{}, &models.StoreProduct{},
		&models.ProductImage{}, &models.SProductImage{}, &models.ProductOption{})
	fmt.Println("👍 Migration complete")
}
