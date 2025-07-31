package grpcclient

import (
	"context"
	"test/domain"
	"test/model"
)

type ProductServiceAdapter struct {
	client *ProductClient
}

func NewProductServiceAdapter(client *ProductClient) domain.ProductService {
	return &ProductServiceAdapter{
		client: client,
	}
}

func (a *ProductServiceAdapter) Create(ctx context.Context, product *model.Product) error {
	return a.client.CreateProduct(ctx, product)
}

func (a *ProductServiceAdapter) GetAll(ctx context.Context) ([]model.Product, error) {
	return a.client.GetProducts(ctx, 1, 100, "", nil)
}

func (a *ProductServiceAdapter) GetByID(ctx context.Context, id uint) (*model.Product, error) {
	return a.client.GetProduct(ctx, id)
}

func (a *ProductServiceAdapter) GetPaginated(ctx context.Context, limit, offset int, query string, filter model.ProductFilter) ([]model.Product, error) {
	page := (offset / limit) + 1
	return a.client.GetProducts(ctx, int(page), limit, query, &filter)
}

func (a *ProductServiceAdapter) Update(ctx context.Context, id uint64, product *model.Product) error {
	return a.client.UpdateProduct(ctx, id, product)
} 