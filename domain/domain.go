package domain

import (
	"context"
)

type Product struct {
	ID    uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title string  `gorm:"type:varchar(255);not null" form:"title" binding:"required" json:"title"`
	Price float64 `gorm:"type:decimal(10,2);default:0" json:"price"`
}

type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
	FetchAll(ctx context.Context) ([]Product, error)
	FindByID(ctx context.Context, id uint) (*Product, error)
	FetchPaginated(ctx context.Context, limit, offset int) ([]Product, error)
	Update(ctx context.Context, id uint, product *Product) error
}

type ProductUsecase interface {
	Create(ctx context.Context, product *Product) error
	FetchAll(ctx context.Context) ([]Product, error)
	FindByID(ctx context.Context, id uint) (*Product, error)
	FetchPaginated(ctx context.Context, limit, offset int) ([]Product, error)
	Update(ctx context.Context, id uint, product *Product) error
}
