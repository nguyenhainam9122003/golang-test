package models

// Product model cho GraphQL service
type Product struct {
	ID                uint64  `json:"id"`
	Name              string  `json:"name"`
	Code              string  `json:"code"`
	ShortCode         *string `json:"short_code"`
	ProductType       string  `json:"product_type"`
	VAT               bool    `json:"vat"`
	VatPercent        int     `json:"vat_percent"`
	VatValue          int64   `json:"vat_value"`
	CategoryID        *uint64 `json:"category_id"`
	ProductGroupID    *uint64 `json:"product_group_id"`
	PlatformID        uint64  `json:"platform_id"`
	CanPreOrder       bool    `json:"can_pre_order"`
	ImageURL          *string `json:"image_url"`
	Note              *string `json:"note"`
	MinOrderQuantity  int64   `json:"min_order_quantity"`
	MedusaID          *string `json:"medusa_id"`
	OriginalCode      *string `json:"original_code"`
	OriginalProductID *uint64 `json:"original_product_id"`
	SearchTextV2      *string `json:"search_text_v2"`
	SearchText        *string `json:"search_text"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}

// CreateProductInput model cho GraphQL input
type CreateProductInput struct {
	Name              string  `json:"name"`
	Code              string  `json:"code"`
	ShortCode         *string `json:"short_code"`
	ProductType       string  `json:"product_type"`
	VAT               bool    `json:"vat"`
	VatPercent        int     `json:"vat_percent"`
	VatValue          int64   `json:"vat_value"`
	CategoryID        *uint64 `json:"category_id"`
	ProductGroupID    *uint64 `json:"product_group_id"`
	PlatformID        uint64  `json:"platform_id"`
	CanPreOrder       bool    `json:"can_pre_order"`
	ImageURL          *string `json:"image_url"`
	Note              *string `json:"note"`
	MinOrderQuantity  int64   `json:"min_order_quantity"`
	MedusaID          *string `json:"medusa_id"`
	OriginalCode      *string `json:"original_code"`
	OriginalProductID *uint64 `json:"original_product_id"`
}

// UpdateProductInput model cho GraphQL update input
type UpdateProductInput struct {
	Name              *string  `json:"name,omitempty"`
	Code              *string  `json:"code,omitempty"`
	ShortCode         *string  `json:"short_code,omitempty"`
	ProductType       *string  `json:"product_type,omitempty"`
	VAT               *bool    `json:"vat,omitempty"`
	VatPercent        *int     `json:"vat_percent,omitempty"`
	VatValue          *int64   `json:"vat_value,omitempty"`
	CategoryID        *uint64  `json:"category_id,omitempty"`
	ProductGroupID    *uint64  `json:"product_group_id,omitempty"`
	PlatformID        *uint64  `json:"platform_id,omitempty"`
	CanPreOrder       *bool    `json:"can_pre_order,omitempty"`
	ImageURL          *string  `json:"image_url,omitempty"`
	Note              *string  `json:"note,omitempty"`
	MinOrderQuantity  *int64   `json:"min_order_quantity,omitempty"`
	MedusaID          *string  `json:"medusa_id,omitempty"`
	OriginalCode      *string  `json:"original_code,omitempty"`
	OriginalProductID *uint64  `json:"original_product_id,omitempty"`
}

// ProductFilter model cho filtering
type ProductFilter struct {
	Category *string  `json:"category,omitempty"`
	MinPrice *float64 `json:"min_price,omitempty"`
	MaxPrice *float64 `json:"max_price,omitempty"`
	InStock  *bool    `json:"in_stock,omitempty"`
}

// ProductPagination model cho pagination
type ProductPagination struct {
	Page  int32      `json:"page"`
	Limit int32      `json:"limit"`
	Items []*Product `json:"items"`
	Total int32      `json:"total"`
}

// MaterialProduct model cho material products
type MaterialProduct struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// APIResponse model cho HTTP API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
} 