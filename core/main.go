package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"test/core/db"
	"test/core/repository"
	myServer "test/core/server"
	"test/core/service"
	protoProduct "test/proto/gen/product"
)

const defaultPort = "50051" // gRPC core port mặc định

func main() {
	// Initialize database
	database := db.Init()

	// Initialize repository and service
	productRepo := repository.NewProductRepository(database)
	productService := service.NewProductService(productRepo)

	// Create gRPC core
	server := grpc.NewServer()
	protoProduct.RegisterProductServiceServer(server, &myServer.ProductServer{
		ProductService: productService,
	})

	// Start core
	lis, err := net.Listen("tcp", ":"+defaultPort)
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	log.Printf("🚀 gRPC core listening on :%s", defaultPort)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}
