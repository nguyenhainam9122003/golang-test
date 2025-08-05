package model

import (
	"errors"
	"test/grpc_service/utils"
	"time"
)

type Product struct {
	ID                uint64         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CreatedAt         time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	Type              string         `gorm:"column:type;type:varchar(50);not null" json:"type"`
	UomID             *uint64        `gorm:"column:uom_id" json:"uom_id"`
	IsActive          bool           `gorm:"column:is_active;type:tinyint(1);default:1;not null" json:"is_active"`
	SellingStatus     SellingStatus  `gorm:"column:selling_status;type:enum('AVAILABLE','CONTACT','UNAVAILABLE');default:'AVAILABLE';not null" json:"selling_status"`
	Name              string         `gorm:"column:name;type:varchar(500);not null" json:"name"`
	Code              string         `gorm:"column:code;type:varchar(100);not null" json:"code"`
	ShortCode         *string        `gorm:"column:short_code;type:varchar(100)" json:"short_code"`
	ProductType       ProductType    `gorm:"column:product_type;type:enum('MATERIAL','SERVICE','VOUCHER','KEY_LICENSE','ACCOUNT');default:'MATERIAL';not null" json:"product_type"`
	VAT               bool           `gorm:"column:vat;type:tinyint(1);default:0;not null" json:"vat"`
	VatType           *VatType       `gorm:"column:vat_type;type:enum('REQUIRED','NOT_REQUIRED','OPTIONAL')" json:"vat_type"`
	Warranty          *uint64        `gorm:"column:warranty" json:"warranty"`
	WarrantyUnit      *WarrantyUnit  `gorm:"column:warranty_unit;type:enum('YEAR','MONTH','DAY')" json:"warranty_unit"`
	VatPercent        int            `gorm:"column:vat_percent;default:0;not null" json:"vat_percent"`
	VatValue          int64          `gorm:"column:vat_value;default:0;not null" json:"vat_value"`
	CategoryID        *uint64        `gorm:"column:category_id" json:"category_id"`
	ProductGroupID    *uint64        `gorm:"column:product_group_id" json:"product_group_id"`
	PlatformID        uint64         `gorm:"column:platform_id;default:1;not null" json:"platform_id"`
	CanPreOrder       bool           `gorm:"column:can_pre_order;type:tinyint(1);default:1;not null" json:"can_pre_order"`
	CopyrightTerm     *uint64        `gorm:"column:copyright_term" json:"copyright_term"`
	CopyrightUnit     *CopyrightUnit `gorm:"column:copyright_unit;type:enum('YEAR','MONTH','DAY','FOREVER')" json:"copyright_unit"`
	ImageURL          *string        `gorm:"column:image_url;type:text" json:"image_url"`
	Note              *string        `gorm:"column:note;type:text" json:"note"`
	MinOrderQuantity  int64          `gorm:"column:min_order_quantity;default:1;not null" json:"min_order_quantity"`
	MedusaID          *string        `gorm:"column:medusa_id;type:varchar(100)" json:"medusa_id"`
	OriginalCode      *string        `gorm:"column:original_code;type:varchar(100)" json:"original_code"`
	OriginalProductID *uint64        `gorm:"column:original_product_id" json:"original_product_id"`

	// GORM ignores generated columns in migrations by default
	SearchTextV2 *string `gorm:"->;column:search_text_v2" json:"search_text_v2"` // generated
	SearchText   *string `gorm:"->;column:search_text" json:"search_text"`       // generated

	ProductPrice *ProductPrice `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"product_price"`
}

func (pt ProductType) IsValid() bool {
	return utils.IsValidEnum(pt, []ProductType{
		ProductTypeMaterial,
		ProductTypeService,
		ProductTypeVoucher,
		ProductTypeKeyLicense,
		ProductTypeAccount,
	})
}

func (ss SellingStatus) IsValid() bool {
	return utils.IsValidEnum(ss, []SellingStatus{
		SellingStatusAvailable,
		SellingStatusUnavailable,
		SellingStatusContact,
	})
}

func (p Product) IsValidVatPercent() bool {
	return p.VatPercent >= 0 && p.VatPercent <= 100
}

func (p *Product) Validate() error {
	if !p.ProductType.IsValid() {
		return errors.New("Invalid product type")
	}
	if !p.SellingStatus.IsValid() {
		return errors.New("Invalid selling status")
	}
	if !p.IsValidVatPercent() {
		return errors.New("VAT percent must be between 0 and 100")
	}
	if !p.ProductPrice.IsValidPrice() {
		return errors.New("All price levels must be greater than 0")
	}
	return nil
}
