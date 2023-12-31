package main

import (
	"AnaWarehouseService/config"
	"AnaWarehouseService/controller"
	"AnaWarehouseService/docs"
	"AnaWarehouseService/repository"
	"AnaWarehouseService/service"
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
	docs.SwaggerInfo.Description = "API for warehouse service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("swagger_host")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define your routes
	router.POST("/warehouse/transfer", productController.StocksTransferRequest)
	router.POST("/warehouse/status", productController.UpdateStatusWarehouse)

	// Start the server
	if err := router.Run(":8084"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
