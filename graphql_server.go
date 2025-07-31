package main

import (
	"log"
	"net/http"
	"os"
	"test/db"
	"test/graph"
	"test/repository"
	"test/service"

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

	database := db.Init()

	productRepo := repository.NewProductRepository(database)
	productService := service.NewProductService(productRepo)

	// Create GraphQL resolver
	resolver := &graph.Resolver{
		ProductService: productService,
	}

	// Create GraphQL server with schema
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
