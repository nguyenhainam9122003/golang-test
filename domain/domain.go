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
	FetchAll(ctx context.Context) ([]Product, error)
	FindByID(ctx context.Context, id uint) (*Product, error)
}

type ProductUsecase interface {
	FetchAll(ctx context.Context) ([]Product, error)
	FindByID(ctx context.Context, id uint) (*Product, error)
}
