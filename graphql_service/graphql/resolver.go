package graphql

import (
	"context"
	"google.golang.org/grpc"
	"test/graphql_service/grpc_client"
	"test/graphql_service/model"
)

type Resolver struct {
	productClient *grpc_client.ProductClient
}

func NewResolver(conn *grpc.ClientConn) *Resolver {
	return &Resolver{
		productClient: grpc_client.NewProductClient(conn),
	}
}

func (r *queryResolver) GetProductByID(ctx context.Context, id string) (*model.MaterialProduct, error) {
	return r.productClient.GetProduct(ctx, id)
}
