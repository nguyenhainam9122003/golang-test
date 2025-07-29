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

func (u *productUsecase) FetchAll(ctx context.Context) ([]domain.Product, error) {
	return u.repo.FetchAll(ctx)
}

func (u *productUsecase) FindByID(ctx context.Context, id uint) (*domain.Product, error) {
	return u.repo.FindByID(ctx, id)
}
