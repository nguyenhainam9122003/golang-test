package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"test/db"
	"test/handler"
	"test/repository"
	"test/service"

	"github.com/gin-gonic/gin"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	database := db.Init()

	productRepo := repository.NewProductRepository(database)
	productService := service.NewProductService(productRepo)

	r := gin.Default()
	handler.NewProductHandler(r, productService)

	// Run server
	if err := r.Run(":" + port); err != nil {
		log.Fatal("failed to start server:", err)
	}

}
