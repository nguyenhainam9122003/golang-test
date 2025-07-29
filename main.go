package main

import (
	"github.com/joho/godotenv"
	"os"
	"test/db"
	"test/handler"
	"test/repository"
	"test/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	database := db.Init()

	productRepo := repository.NewProductRepository(database)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUsecase)

	r := gin.Default()
	api := r.Group("/api")

	api.POST("/products", productHandler.CreateProduct)

	api.GET("/products/paginated", productHandler.GetPaginatedProducts)

	api.GET("/products", productHandler.GetAllProducts)

	api.GET("/products/:id", productHandler.GetProductByID)

	api.PUT("/products/:id", productHandler.UpdateProduct)

	r.Run(":" + port)
}
