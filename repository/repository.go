package repository

import (
	"context"
	"gorm.io/gorm"
	"test/domain"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) FetchAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.WithContext(ctx).Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(ctx context.Context, id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.WithContext(ctx).First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
