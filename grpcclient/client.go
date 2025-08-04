package grpcclient

import (
	"context"
	"test/model"
	grpcproduct "test/proto/product"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClient struct {
	client grpcproduct.ProductServiceClient
}

func NewProductClient(serverAddr string) (*ProductClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	
	client := grpcproduct.NewProductServiceClient(conn)
	return &ProductClient{
		client: client,
	}, nil
}

func (c *ProductClient) CreateProduct(ctx context.Context, product *model.Product) error {
	protoProduct := convertDomainToProtoProduct(*product)
	
	req := &grpcproduct.CreateProductRequest{
		Product: protoProduct,
	}
	
	_, err := c.client.CreateProduct(ctx, req)
	return err
}

func (c *ProductClient) GetProduct(ctx context.Context, id uint) (*model.Product, error) {
	req := &grpcproduct.GetProductRequest{
		Id: uint64(id),
	}
	
	resp, err := c.client.GetProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	
	domainProduct := convertProtoToDomainProduct(resp.Product)
	return &domainProduct, nil
}

func (c *ProductClient) GetProducts(ctx context.Context, page, limit int, query string, filter *model.ProductFilter) ([]model.Product, error) {
	req := &grpcproduct.GetProductsRequest{
		Page:  int32(page),
		Limit: int32(limit),
		Query: &query,
	}
	
	if filter != nil {
		protoFilter := convertDomainToProtoFilter(filter)
		req.Filter = &protoFilter
	}
	
	resp, err := c.client.GetProducts(ctx, req)
	if err != nil {
		return nil, err
	}
	
	var domainProducts []model.Product
	for _, protoProduct := range resp.Products {
		domainProduct := convertProtoToDomainProduct(protoProduct)
		domainProducts = append(domainProducts, domainProduct)
	}
	
	return domainProducts, nil
}

func (c *ProductClient) UpdateProduct(ctx context.Context, id uint64, product *model.Product) error {
	protoProduct := convertDomainToProtoProduct(*product)
	
	req := &grpcproduct.UpdateProductRequest{
		Id:      id,
		Product: protoProduct,
	}
	
	_, err := c.client.UpdateProduct(ctx, req)
	return err
}

func (c *ProductClient) DeleteProduct(ctx context.Context, id uint64) error {
	req := &grpcproduct.DeleteProductRequest{
		Id: id,
	}
	
	_, err := c.client.DeleteProduct(ctx, req)
	return err
} 