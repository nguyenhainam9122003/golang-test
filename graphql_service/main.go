package main

import (
	"log"
	"net/http"
	"os"
	"test/graphql_service/graphql"
	"test/graphql_service/http_client"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "8081" // GraphQL service chạy trên port khác

func main() {
	godotenv.Load()
	port := os.Getenv("GRAPHQL_PORT")
	if port == "" {
		port = defaultPort
	}

	// Khởi tạo HTTP client để gọi API
	apiBaseURL := os.Getenv("API_BASE_URL")
	if apiBaseURL == "" {
		apiBaseURL = "http://localhost:8080" // Default API server
	}

	httpClient := http_client.NewProductHTTPClient(apiBaseURL)

	// Tạo GraphQL resolver với HTTP client
	resolver := &graphql.Resolver{
		ProductHTTPClient: httpClient,
	}

	// Tạo GraphQL server
	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver}))

	// Setup routes
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("GraphQL service starting on http://localhost:%s", port)
	log.Printf("GraphQL playground available at http://localhost:%s", port)
	log.Printf("API base URL: %s", apiBaseURL)
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
} 