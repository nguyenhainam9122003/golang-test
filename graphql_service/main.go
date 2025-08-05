package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"test/graphql_service/graphql"
	"test/graphql_service/graphql/generated"
)

const defaultPort = "8081" // GraphQL service chạy trên port khác

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	resolver := graphql.NewResolver(conn)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("🚀 GraphQL server listening on :%s", port)
	log.Printf("📊 Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Printf("🔗 GraphQL → gRPC → Service → Database")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
