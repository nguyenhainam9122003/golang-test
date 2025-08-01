package service

import (
	"context"
	"test/domain"
	"test/model"
)

type productService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) domain.ProductService {
	return &productService{repo}
}

func (s *productService) Create(ctx context.Context, product *model.Product) error {
	return s.repo.Create(ctx, product)
}

func (s *productService) GetAll(ctx context.Context) ([]model.Product, error) {
	return s.repo.FetchAll(ctx)
}

func (s *productService) GetByID(ctx context.Context, id uint) (*model.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *productService) GetMaterialProduct(ctx context.Context) ([]model.Product, error) {
	return s.repo.GetMaterialProduct(ctx)
}

func (s *productService) GetPaginated(ctx context.Context, limit, offset int, query string, filter model.ProductFilter) ([]model.Product, error) {
	return s.repo.FetchPaginated(ctx, limit, offset, query, filter)
}

func (s *productService) Update(ctx context.Context, id uint64, product *model.Product) error {
	return s.repo.Update(ctx, id, product)
}
