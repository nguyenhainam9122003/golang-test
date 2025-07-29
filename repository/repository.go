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

func (r *productRepository) Create(ctx context.Context, product *domain.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
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

func (r *productRepository) FetchPaginated(ctx context.Context, limit, offset int) ([]domain.Product, error) {
	var products []domain.Product
	result := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&products)
	return products, result.Error
}

func (r *productRepository) Update(ctx context.Context, id uint, product *domain.Product) error {
	var existing domain.Product
	if err := r.db.WithContext(ctx).First(&existing, id).Error; err != nil {
		return err
	}

	product.ID = id // ensure ID is retained
	return r.db.WithContext(ctx).Save(product).Error
}
