package domain

import (
	"context"
	"test/model"
)

type ProductRepository interface {
	Create(ctx context.Context, product *model.Product) error
	FetchAll(ctx context.Context) ([]model.Product, error)
	FindByID(ctx context.Context, id uint) (*model.Product, error)
	GetMaterialProduct(ctx context.Context) ([]model.Product, error)
	FetchPaginated(ctx context.Context, limit, offset int, query string, filter model.ProductFilter) ([]model.Product, error)
	Update(ctx context.Context, id uint64, product *model.Product) error
}

type ProductService interface {
	Create(ctx context.Context, product *model.Product) error
	GetAll(ctx context.Context) ([]model.Product, error)
	GetByID(ctx context.Context, id uint) (*model.Product, error)
	GetMaterialProduct(ctx context.Context) ([]model.Product, error)
	GetPaginated(ctx context.Context, limit, offset int, query string, filter model.ProductFilter) ([]model.Product, error)
	Update(ctx context.Context, id uint64, product *model.Product) error
}
