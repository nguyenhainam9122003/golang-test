package model

import (
	"time"
)

type ProductPrice struct {
	ID         uint64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ProductID  *uint64 `gorm:"column:product_id;uniqueIndex;index" json:"product_id"`
	PlatformID uint64  `gorm:"column:platform_id;default:1;not null;index" json:"platform_id"`

	Level1Price int64 `gorm:"column:level1_price;not null;default:0" json:"level1_price"`
	Level2Price int64 `gorm:"column:level2_price;not null;default:0" json:"level2_price"`
	Level3Price int64 `gorm:"column:level3_price;not null;default:0" json:"level3_price"`
	Level4Price int64 `gorm:"column:level4_price;not null;default:0" json:"level4_price"`

	Level1Quantity *int64 `gorm:"column:level1_quantity" json:"level1_quantity"`
	Level1Discount *int64 `gorm:"column:level1_discount" json:"level1_discount"`
	Level2Quantity *int64 `gorm:"column:level2_quantity" json:"level2_quantity"`
	Level2Discount *int64 `gorm:"column:level2_discount" json:"level2_discount"`
	Level3Quantity *int64 `gorm:"column:level3_quantity" json:"level3_quantity"`
	Level3Discount *int64 `gorm:"column:level3_discount" json:"level3_discount"`
	Level4Quantity *int64 `gorm:"column:level4_quantity" json:"level4_quantity"`
	Level4Discount *int64 `gorm:"column:level4_discount" json:"level4_discount"`

	PriceHasVAT *int64 `gorm:"column:price_has_vat" json:"price_has_vat"`

	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

	// Quan hệ
	Product *Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}

func (pp ProductPrice) IsValidPrice() bool {
	return pp.Level1Price > 0 &&
		pp.Level2Price > 0 &&
		pp.Level3Price > 0 &&
		pp.Level4Price > 0
}
