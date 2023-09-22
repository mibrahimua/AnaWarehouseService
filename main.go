package main

import (
	"AnaProductService/config"
	"AnaProductService/controller"
	"AnaProductService/docs"
	"AnaProductService/repository"
	"AnaProductService/service"
	_ "fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

func main() {
	db := config.GetDB()

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	docs.SwaggerInfo.Title = "Ana Store - Environment: "
	docs.SwaggerInfo.Description = "API for product service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("swagger_host")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define your routes
	router.POST("/product", productController.GetListProductByName)

	// Start the server
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
