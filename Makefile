.PHONY: proto build-grpc build-graphql run-grpc run-graphql run-all clean

# Generate proto files
proto:
	protoc --experimental_allow_proto3_optional --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/product.proto

# Build gRPC server
build-grpc:
	go build -o grpc_server grpc/server.go

# Build GraphQL server
build-graphql:
	go build -o graphql_server graphql_server.go

# Build GraphQL server with gRPC client
build-graphql-grpc:
	go build -o graphql_server_grpc graphql_server_grpc.go

# Run gRPC server
run-grpc:
	./grpc_server

# Run GraphQL server (direct service)
run-graphql:
	./graphql_server

# Run GraphQL server (with gRPC client)
run-graphql-grpc:
	./graphql_server_grpc

# Run both servers
run-all:
	@echo "Starting gRPC server..."
	@make run-grpc &
	@sleep 2
	@echo "Starting GraphQL server..."
	@make run-graphql-grpc

# Clean build files
clean:
	rm -f grpc_server graphql_server graphql_server_grpc
	rm -f proto/*.pb.go

# Install dependencies
deps:
	go mod tidy
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Setup everything
setup: deps proto build-grpc build-graphql build-graphql-grpc 