package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"test/grpc_service/db"
	"test/grpc_service/repository"
	myServer "test/grpc_service/server"
	"test/grpc_service/service"
	protoProduct "test/proto/gen/product"
)

const defaultPort = "50051" // gRPC grpc_server port mặc định

func main() {
	// Initialize database
	database := db.Init()

	// Initialize repository and service
	productRepo := repository.NewProductRepository(database)
	productService := service.NewProductService(productRepo)

	// Create gRPC grpc_server
	server := grpc.NewServer()
	protoProduct.RegisterProductServiceServer(server, &myServer.ProductServer{
		ProductService: productService,
	})

	// Start grpc_server
	lis, err := net.Listen("tcp", ":"+defaultPort)
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	log.Printf("🚀 gRPC grpc_server listening on :%s", defaultPort)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}
