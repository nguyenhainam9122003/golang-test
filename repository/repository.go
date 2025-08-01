package repository

import (
	"context"
	"gorm.io/gorm"
	"test/domain"
	"test/model"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(ctx context.Context, product *model.Product) error {
	tx := r.db.WithContext(ctx)

	var price *model.ProductPrice = product.ProductPrice
	product.ProductPrice = nil

	// Bước 1: Tạo Product trước
	if err := tx.Create(product).Error; err != nil {
		return err
	}

	// Bước 2: Nếu có ProductPrice thì gán ProductID và tạo tiếp
	if price != nil {
		price.ProductID = &product.ID
		if err := tx.Create(price).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *productRepository) FetchAll(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	err := r.db.WithContext(ctx).Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(ctx context.Context, id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.WithContext(ctx).
		Preload("ProductPrice").
		First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) FetchPaginated(ctx context.Context, limit, offset int, query string, filter model.ProductFilter) ([]model.Product, error) {
	var products []model.Product

	tx := r.db.WithContext(ctx).Preload("ProductPrice").Limit(limit).Offset(offset)
	if query != "" {
		// Tìm theo tên hoặc mô tả (tuỳ theo trường bạn muốn)
		tx = tx.Where("name LIKE ? OR code LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	if filter.ProductType != nil {
		tx = tx.Where("product_type = ?", *filter.ProductType)
	}

	if filter.SellingStatus != nil {
		tx = tx.Where("selling_status = ?", *filter.SellingStatus)
	}

	err := tx.Find(&products).Error
	return products, err
}

func (r *productRepository) Update(ctx context.Context, id uint64, product *model.Product) error {
	var existing model.Product
	if err := r.db.WithContext(ctx).First(&existing, id).Error; err != nil {
		return err
	}
	product.ID = id
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Cập nhật Product
		if err := tx.Model(&model.Product{}).
			Where("id = ?", id).
			Updates(product).Error; err != nil {
			return err
		}

		// Cập nhật ProductPrice tương ứng
		if err := tx.Model(&model.ProductPrice{}).
			Where("product_id = ?", id).
			Updates(product.ProductPrice).Error; err != nil {
			return err
		}

		return nil
	})
}
