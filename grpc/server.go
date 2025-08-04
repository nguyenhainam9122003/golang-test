package main

import (
	"log"
	"net"
	"test/db"
	grpcproduct "test/proto/product"
	"test/repository"
	"test/service"

	"google.golang.org/grpc"
)

func main() {
	log.Printf("🚀 Starting gRPC server...")
	
	// Initialize database
	database := db.Init()
	log.Printf("✅ Database connected")
	
	// Initialize repository and service
	productRepo := repository.NewProductRepository(database)
	productService := service.NewProductService(productRepo)
	log.Printf("✅ Repository and Service initialized")
	
	// Create gRPC server
	server := grpc.NewServer()
	grpcproduct.RegisterProductServiceServer(server, &productServer{
		productService: productService,
	})
	log.Printf("✅ gRPC server registered")
	
	// Start server
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}
	
	log.Printf("🚀 gRPC server listening on :9090")
	log.Printf("🔗 Ready to receive requests from GraphQL server")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
} 