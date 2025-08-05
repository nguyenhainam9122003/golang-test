package model

import "fmt"

type ProductFilter struct {
	ProductType   *string `form:"product_type"`
	SellingStatus *string `form:"selling_status"`
}

func (f *ProductFilter) Validate() error {
	if f.ProductType != nil && *f.ProductType != "" {
		pt := ProductType(*f.ProductType)
		if !pt.IsValid() {
			return fmt.Errorf("Loại sản phẩm không hợp lệ: %s", *f.ProductType)
		}
	}

	if f.SellingStatus != nil && *f.SellingStatus != "" {
		ss := SellingStatus(*f.SellingStatus)
		if !ss.IsValid() {
			return fmt.Errorf("Trạng thái bán không hợp lệ: %s", *f.SellingStatus)
		}
	}
	return nil
}
