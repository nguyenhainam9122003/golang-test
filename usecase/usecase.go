package usecase

import (
	"context"
	"test/domain"
)

type productUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(repo domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{repo}
}

func (u *productUsecase) Create(ctx context.Context, product *domain.Product) error {
	return u.repo.Create(ctx, product)
}

func (u *productUsecase) FetchAll(ctx context.Context) ([]domain.Product, error) {
	return u.repo.FetchAll(ctx)
}

func (u *productUsecase) FindByID(ctx context.Context, id uint) (*domain.Product, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *productUsecase) FetchPaginated(ctx context.Context, limit, offset int) ([]domain.Product, error) {
	return u.repo.FetchPaginated(ctx, limit, offset)
}

func (u *productUsecase) Update(ctx context.Context, id uint, product *domain.Product) error {
	return u.repo.Update(ctx, id, product)
}
