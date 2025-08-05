package grpc_client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"test/graphql_service/model"
	pb "test/proto/gen/product"
)

type ProductClient struct {
	client pb.ProductServiceClient
}

func NewProductClient(conn *grpc.ClientConn) *ProductClient {
	return &ProductClient{
		client: pb.NewProductServiceClient(conn),
	}
}

func (pc *ProductClient) GetProduct(ctx context.Context, id string) (*model.MaterialProduct, error) {
	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	resp, err := pc.client.GetProduct(ctx, &pb.GetProductRequest{Id: parsedID})
	if err != nil {
		log.Printf("gRPC error: %v", err)
		return nil, err
	}

	p := resp.GetProduct()
	return &model.MaterialProduct{
		ID:   strconv.FormatUint(p.Id, 10),
		Name: p.Name,
		Code: p.Code,
		// Map các trường khác nếu có
	}, nil
}
