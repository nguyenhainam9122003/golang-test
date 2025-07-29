package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"test/db"
	"test/handler"
	"test/repository"
	"test/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Trước khi handler chạy
		path := c.Request.URL.Path
		method := c.Request.Method
		fmt.Println("Before request:", method, path)

		c.Next() // gọi handler tiếp theo

		// Sau khi handler chạy
		fmt.Println("After request:", time.Since(t))
	}
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	database := db.Init()

	productRepo := repository.NewProductRepository(database)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUsecase)

	r := gin.Default()
	r.Use(Logger())
	api := r.Group("/api")

	api.POST("/products", productHandler.CreateProduct)

	api.GET("/products/paginated", productHandler.GetPaginatedProducts)

	api.GET("/products", productHandler.GetAllProducts)

	api.GET("/products/:id", productHandler.GetProductByID)

	api.PUT("/products/:id", productHandler.UpdateProduct)

	r.Run(":" + port)
}
