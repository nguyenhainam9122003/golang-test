package graphql

import (
	"google.golang.org/grpc"
	"test/graphql_service/grpc_client"
)

type Resolver struct {
	productClient *grpc_client.ProductClient
}

func NewResolver(conn *grpc.ClientConn) *Resolver {
	return &Resolver{
		productClient: grpc_client.NewProductClient(conn),
	}
}


