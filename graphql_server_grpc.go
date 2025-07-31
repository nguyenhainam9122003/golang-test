package main

import (
	"log"
	"net/http"
	"os"
	"test/graph"
	"test/grpcclient"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize gRPC client
	grpcClient, err := grpcclient.NewProductClient("localhost:9090")
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	// Create adapter to implement domain.ProductService interface
	productService := grpcclient.NewProductServiceAdapter(grpcClient)

	// Create GraphQL resolver with gRPC client
	resolver := &graph.Resolver{
		ProductService: productService,
	}

	// Create GraphQL server with schema
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("GraphQL server listening on :%s", port)
	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
} 