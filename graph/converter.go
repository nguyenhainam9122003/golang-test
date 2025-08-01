package graph

import (
	"fmt"
	"strconv"
	"test/graph/model"
	domainmodel "test/model"
)

func ConvertToGraphQLProduct(p domainmodel.Product) *model.Product {
	idStr := fmt.Sprintf("%d", p.ID)
	var sellingStatus model.SellingStatus
	if p.SellingStatus != "" {
		status := ConvertSellingStatus(p.SellingStatus)
		if status != nil {
			sellingStatus = *status
		}
	}
	var warranty *int32
	if p.Warranty != nil {
		w := int32(*p.Warranty)
		warranty = &w
	}
	var warrantyUnit *model.WarrantyUnit
	if p.WarrantyUnit != nil {
		unit := ConvertWarrantyUnit(p.WarrantyUnit)
		warrantyUnit = unit
	}
	var copyrightTerm *int32
	if p.CopyrightTerm != nil {
		ct := int32(*p.CopyrightTerm)
		copyrightTerm = &ct
	}
	var copyrightUnit *model.CopyrightUnit
	if p.CopyrightUnit != nil {
		unit := ConvertCopyrightUnit(p.CopyrightUnit)
		copyrightUnit = unit
	}
	return &model.Product{
		ID:                idStr,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
		Type:              p.Type,
		UomID:             ConvertUint64ToString(p.UomID),
		IsActive:          p.IsActive,
		SellingStatus:     sellingStatus,
		Name:              p.Name,
		Code:              p.Code,
		ShortCode:         p.ShortCode,
		ProductType:       ConvertProductType(p.ProductType),
		Vat:               p.VAT,
		VatType:           ConvertVatType(p.VatType),
		Warranty:          warranty,
		WarrantyUnit:      warrantyUnit,
		VatPercent:        int32(p.VatPercent),
		VatValue:          int32(p.VatValue),
		CategoryID:        ConvertUint64ToString(p.CategoryID),
		ProductGroupID:    ConvertUint64ToString(p.ProductGroupID),
		PlatformID:        fmt.Sprintf("%d", p.PlatformID),
		CanPreOrder:       p.CanPreOrder,
		CopyrightTerm:     copyrightTerm,
		CopyrightUnit:     copyrightUnit,
		ImageURL:          p.ImageURL,
		Note:              p.Note,
		MinOrderQuantity:  int32(p.MinOrderQuantity),
		MedusaID:          p.MedusaID,
		OriginalCode:      p.OriginalCode,
		OriginalProductID: ConvertUint64ToString(p.OriginalProductID),
		SearchTextV2:      p.SearchTextV2,
		SearchText:        p.SearchText,
		ProductPrice:      ConvertToGraphQLProductPrice(p.ProductPrice),
	}
}
func ConvertToGraphQLProductMaterial(p domainmodel.Product) *model.MaterialProduct {
	idStr := fmt.Sprintf("%d", p.ID)
	var sellingStatus model.SellingStatus
	if p.SellingStatus != "" {
		status := ConvertSellingStatus(p.SellingStatus)
		if status != nil {
			sellingStatus = *status
		}
	}

	return &model.MaterialProduct{
		ID:            idStr,
		Type:          p.Type,
		IsActive:      p.IsActive,
		SellingStatus: sellingStatus,
		Name:          p.Name,
		Code:          p.Code,
		ShortCode:     p.ShortCode,
		ProductType:   ConvertProductType(p.ProductType),
	}
}

func ConvertToProductModel(input model.CreateProductInput) domainmodel.Product {
	var isActive bool = true
	if input.IsActive != nil {
		isActive = *input.IsActive
	}
	var canPreOrder bool = true
	if input.CanPreOrder != nil {
		canPreOrder = *input.CanPreOrder
	}
	var vatPercent int = 0
	if input.VatPercent != nil {
		vatPercent = int(*input.VatPercent)
	}
	var vatValue int64 = 0
	if input.VatValue != nil {
		vatValue = int64(*input.VatValue)
	}
	var minOrderQuantity int64 = 1
	if input.MinOrderQuantity != nil {
		minOrderQuantity = int64(*input.MinOrderQuantity)
	}
	var platformID uint64 = 1
	if input.PlatformID != nil {
		platformID = *ConvertStringToUint64(input.PlatformID)
	}
	var warranty *uint64
	if input.Warranty != nil {
		w := uint64(*input.Warranty)
		warranty = &w
	}
	var copyrightTerm *uint64
	if input.CopyrightTerm != nil {
		ct := uint64(*input.CopyrightTerm)
		copyrightTerm = &ct
	}
	var vat bool = false
	if input.Vat != nil {
		vat = *input.Vat
	}
	return domainmodel.Product{
		Type:              input.Type,
		UomID:             ConvertStringToUint64(input.UomID),
		IsActive:          isActive,
		SellingStatus:     ConvertSellingStatusToDomain(input.SellingStatus),
		Name:              input.Name,
		Code:              input.Code,
		ShortCode:         input.ShortCode,
		ProductType:       ConvertProductTypeToDomain(&input.ProductType),
		VAT:               vat,
		VatType:           ConvertVatTypeToDomain(input.VatType),
		Warranty:          warranty,
		WarrantyUnit:      ConvertWarrantyUnitToDomain(input.WarrantyUnit),
		VatPercent:        vatPercent,
		VatValue:          vatValue,
		CategoryID:        ConvertStringToUint64(input.CategoryID),
		ProductGroupID:    ConvertStringToUint64(input.ProductGroupID),
		PlatformID:        platformID,
		CanPreOrder:       canPreOrder,
		CopyrightTerm:     copyrightTerm,
		CopyrightUnit:     ConvertCopyrightUnitToDomain(input.CopyrightUnit),
		ImageURL:          input.ImageURL,
		Note:              input.Note,
		MinOrderQuantity:  minOrderQuantity,
		MedusaID:          input.MedusaID,
		OriginalCode:      input.OriginalCode,
		OriginalProductID: ConvertStringToUint64(input.OriginalProductID),
		ProductPrice:      ConvertToProductPriceModel(input.ProductPrice),
	}
}
func ConvertToProductModelFromUpdate(input model.UpdateProductInput) domainmodel.Product {
	product := domainmodel.Product{}
	if input.Type != nil {
		product.Type = *input.Type
	}
	if input.UomID != nil {
		product.UomID = ConvertStringToUint64(input.UomID)
	}
	if input.IsActive != nil {
		product.IsActive = *input.IsActive
	}
	if input.SellingStatus != nil {
		product.SellingStatus = ConvertSellingStatusToDomain(input.SellingStatus)
	}
	if input.Name != nil {
		product.Name = *input.Name
	}
	if input.Code != nil {
		product.Code = *input.Code
	}
	if input.ShortCode != nil {
		product.ShortCode = input.ShortCode
	}
	if input.ProductType != nil {
		product.ProductType = ConvertProductTypeToDomain(input.ProductType)
	}
	if input.Vat != nil {
		product.VAT = *input.Vat
	}
	if input.VatType != nil {
		product.VatType = ConvertVatTypeToDomain(input.VatType)
	}
	if input.Warranty != nil {
		w := uint64(*input.Warranty)
		product.Warranty = &w
	}
	if input.WarrantyUnit != nil {
		product.WarrantyUnit = ConvertWarrantyUnitToDomain(input.WarrantyUnit)
	}
	if input.VatPercent != nil {
		product.VatPercent = int(*input.VatPercent)
	}
	if input.VatValue != nil {
		product.VatValue = int64(*input.VatValue)
	}
	if input.CategoryID != nil {
		product.CategoryID = ConvertStringToUint64(input.CategoryID)
	}
	if input.ProductGroupID != nil {
		product.ProductGroupID = ConvertStringToUint64(input.ProductGroupID)
	}
	if input.PlatformID != nil {
		product.PlatformID = *ConvertStringToUint64(input.PlatformID)
	}
	if input.CanPreOrder != nil {
		product.CanPreOrder = *input.CanPreOrder
	}
	if input.CopyrightTerm != nil {
		ct := uint64(*input.CopyrightTerm)
		product.CopyrightTerm = &ct
	}
	if input.CopyrightUnit != nil {
		product.CopyrightUnit = ConvertCopyrightUnitToDomain(input.CopyrightUnit)
	}
	if input.ImageURL != nil {
		product.ImageURL = input.ImageURL
	}
	if input.Note != nil {
		product.Note = input.Note
	}
	if input.MinOrderQuantity != nil {
		product.MinOrderQuantity = int64(*input.MinOrderQuantity)
	}
	if input.MedusaID != nil {
		product.MedusaID = input.MedusaID
	}
	if input.OriginalCode != nil {
		product.OriginalCode = input.OriginalCode
	}
	if input.OriginalProductID != nil {
		product.OriginalProductID = ConvertStringToUint64(input.OriginalProductID)
	}
	if input.ProductPrice != nil {
		product.ProductPrice = ConvertToProductPriceModelFromUpdate(*input.ProductPrice)
	}
	return product
}
func ConvertToProductFilter(filter *model.ProductFilter) *domainmodel.ProductFilter {
	if filter == nil {
		return nil
	}
	var pt, ss *string
	if filter.ProductType != nil {
		s := string(*filter.ProductType)
		pt = &s
	}
	if filter.SellingStatus != nil {
		s := string(*filter.SellingStatus)
		ss = &s
	}
	return &domainmodel.ProductFilter{
		ProductType:   pt,
		SellingStatus: ss,
	}
}
func ConvertUint64ToString(id *uint64) *string {
	if id == nil {
		return nil
	}
	str := fmt.Sprintf("%d", *id)
	return &str
}
func ConvertStringToUint64(id *string) *uint64 {
	if id == nil {
		return nil
	}
	val, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return nil
	}
	return &val
}
func ConvertInt64ToInt(val *int64) *int {
	if val == nil {
		return nil
	}
	intVal := int(*val)
	return &intVal
}
func ConvertIntToInt64(val *int) *int64 {
	if val == nil {
		return nil
	}
	int64Val := int64(*val)
	return &int64Val
}
func ConvertSellingStatus(status domainmodel.SellingStatus) *model.SellingStatus {
	switch string(status) {
	case string(domainmodel.SellingStatusAvailable):
		val := model.SellingStatusAvailable
		return &val
	case string(domainmodel.SellingStatusContact):
		val := model.SellingStatusContact
		return &val
	case string(domainmodel.SellingStatusUnavailable):
		val := model.SellingStatusUnavailable
		return &val
	default:
		return nil
	}
}
func ConvertSellingStatusToDomain(status *model.SellingStatus) domainmodel.SellingStatus {
	if status == nil {
		return domainmodel.SellingStatusAvailable
	}
	switch *status {
	case model.SellingStatusAvailable:
		return domainmodel.SellingStatusAvailable
	case model.SellingStatusContact:
		return domainmodel.SellingStatusContact
	case model.SellingStatusUnavailable:
		return domainmodel.SellingStatusUnavailable
	default:
		return domainmodel.SellingStatusAvailable
	}
}
func ConvertProductType(productType domainmodel.ProductType) model.ProductType {
	switch string(productType) {
	case string(domainmodel.ProductTypeMaterial):
		return model.ProductTypeMaterial
	case string(domainmodel.ProductTypeService):
		return model.ProductTypeService
	case string(domainmodel.ProductTypeVoucher):
		return model.ProductTypeVoucher
	case string(domainmodel.ProductTypeKeyLicense):
		return model.ProductTypeKeyLicense
	case string(domainmodel.ProductTypeAccount):
		return model.ProductTypeAccount
	default:
		return model.ProductTypeMaterial
	}
}
func ConvertProductTypeToDomain(productType *model.ProductType) domainmodel.ProductType {
	if productType == nil {
		return domainmodel.ProductTypeMaterial
	}
	switch *productType {
	case model.ProductTypeMaterial:
		return domainmodel.ProductTypeMaterial
	case model.ProductTypeService:
		return domainmodel.ProductTypeService
	case model.ProductTypeVoucher:
		return domainmodel.ProductTypeVoucher
	case model.ProductTypeKeyLicense:
		return domainmodel.ProductTypeKeyLicense
	case model.ProductTypeAccount:
		return domainmodel.ProductTypeAccount
	default:
		return domainmodel.ProductTypeMaterial
	}
}
func ConvertVatType(vatType *domainmodel.VatType) *model.VatType {
	if vatType == nil {
		return nil
	}
	switch *vatType {
	case domainmodel.VatTypeRequired:
		val := model.VatTypeRequired
		return &val
	case domainmodel.VatTypeNotRequired:
		val := model.VatTypeNotRequired
		return &val
	case domainmodel.VatTypeOptional:
		val := model.VatTypeOptional
		return &val
	default:
		return nil
	}
}
func ConvertVatTypeToDomain(vatType *model.VatType) *domainmodel.VatType {
	if vatType == nil {
		return nil
	}
	switch *vatType {
	case model.VatTypeRequired:
		val := domainmodel.VatTypeRequired
		return &val
	case model.VatTypeNotRequired:
		val := domainmodel.VatTypeNotRequired
		return &val
	case model.VatTypeOptional:
		val := domainmodel.VatTypeOptional
		return &val
	default:
		return nil
	}
}
func ConvertWarrantyUnit(warrantyUnit *domainmodel.WarrantyUnit) *model.WarrantyUnit {
	if warrantyUnit == nil {
		return nil
	}
	switch *warrantyUnit {
	case domainmodel.WarrantyUnitYear:
		val := model.WarrantyUnitYear
		return &val
	case domainmodel.WarrantyUnitMonth:
		val := model.WarrantyUnitMonth
		return &val
	case domainmodel.WarrantyUnitDay:
		val := model.WarrantyUnitDay
		return &val
	default:
		return nil
	}
}
func ConvertWarrantyUnitToDomain(warrantyUnit *model.WarrantyUnit) *domainmodel.WarrantyUnit {
	if warrantyUnit == nil {
		return nil
	}
	switch *warrantyUnit {
	case model.WarrantyUnitYear:
		val := domainmodel.WarrantyUnitYear
		return &val
	case model.WarrantyUnitMonth:
		val := domainmodel.WarrantyUnitMonth
		return &val
	case model.WarrantyUnitDay:
		val := domainmodel.WarrantyUnitDay
		return &val
	default:
		return nil
	}
}
func ConvertCopyrightUnit(copyrightUnit *domainmodel.CopyrightUnit) *model.CopyrightUnit {
	if copyrightUnit == nil {
		return nil
	}
	switch *copyrightUnit {
	case domainmodel.CopyrightUnitYear:
		val := model.CopyrightUnitYear
		return &val
	case domainmodel.CopyrightUnitMonth:
		val := model.CopyrightUnitMonth
		return &val
	case domainmodel.CopyrightUnitDay:
		val := model.CopyrightUnitDay
		return &val
	case domainmodel.CopyrightUnitForever:
		val := model.CopyrightUnitForever
		return &val
	default:
		return nil
	}
}
func ConvertCopyrightUnitToDomain(copyrightUnit *model.CopyrightUnit) *domainmodel.CopyrightUnit {
	if copyrightUnit == nil {
		return nil
	}
	switch *copyrightUnit {
	case model.CopyrightUnitYear:
		val := domainmodel.CopyrightUnitYear
		return &val
	case model.CopyrightUnitMonth:
		val := domainmodel.CopyrightUnitMonth
		return &val
	case model.CopyrightUnitDay:
		val := domainmodel.CopyrightUnitDay
		return &val
	case model.CopyrightUnitForever:
		val := domainmodel.CopyrightUnitForever
		return &val
	default:
		return nil
	}
}
func ConvertToGraphQLProductPrice(price *domainmodel.ProductPrice) *model.ProductPrice {
	if price == nil {
		return nil
	}
	idStr := fmt.Sprintf("%d", price.ID)
	return &model.ProductPrice{
		ID:             idStr,
		ProductID:      ConvertUint64ToString(price.ProductID),
		PlatformID:     fmt.Sprintf("%d", price.PlatformID),
		Level1Price:    int32(price.Level1Price),
		Level2Price:    int32(price.Level2Price),
		Level3Price:    int32(price.Level3Price),
		Level4Price:    int32(price.Level4Price),
		Level1Quantity: ConvertInt64ToInt32(price.Level1Quantity),
		Level1Discount: ConvertInt64ToInt32(price.Level1Discount),
		Level2Quantity: ConvertInt64ToInt32(price.Level2Quantity),
		Level2Discount: ConvertInt64ToInt32(price.Level2Discount),
		Level3Quantity: ConvertInt64ToInt32(price.Level3Quantity),
		Level3Discount: ConvertInt64ToInt32(price.Level3Discount),
		Level4Quantity: ConvertInt64ToInt32(price.Level4Quantity),
		Level4Discount: ConvertInt64ToInt32(price.Level4Discount),
		PriceHasVat:    ConvertInt64ToInt32(price.PriceHasVAT),
		CreatedAt:      price.CreatedAt,
		UpdatedAt:      price.UpdatedAt,
	}
}
func ConvertToProductPriceModel(input *model.CreateProductPriceInput) *domainmodel.ProductPrice {
	if input == nil {
		return nil
	}
	return &domainmodel.ProductPrice{
		Level1Price:    int64(input.Level1Price),
		Level2Price:    int64(input.Level2Price),
		Level3Price:    int64(input.Level3Price),
		Level4Price:    int64(input.Level4Price),
		Level1Quantity: ConvertInt32ToInt64(input.Level1Quantity),
		Level1Discount: ConvertInt32ToInt64(input.Level1Discount),
		Level2Quantity: ConvertInt32ToInt64(input.Level2Quantity),
		Level2Discount: ConvertInt32ToInt64(input.Level2Discount),
		Level3Quantity: ConvertInt32ToInt64(input.Level3Quantity),
		Level3Discount: ConvertInt32ToInt64(input.Level3Discount),
		Level4Quantity: ConvertInt32ToInt64(input.Level4Quantity),
		Level4Discount: ConvertInt32ToInt64(input.Level4Discount),
		PriceHasVAT:    ConvertInt32ToInt64(input.PriceHasVat),
	}
}
func ConvertToProductPriceModelFromUpdate(input model.UpdateProductPriceInput) *domainmodel.ProductPrice {
	price := &domainmodel.ProductPrice{}
	if input.Level1Price != nil {
		price.Level1Price = int64(*input.Level1Price)
	}
	if input.Level2Price != nil {
		price.Level2Price = int64(*input.Level2Price)
	}
	if input.Level3Price != nil {
		price.Level3Price = int64(*input.Level3Price)
	}
	if input.Level4Price != nil {
		price.Level4Price = int64(*input.Level4Price)
	}
	if input.Level1Quantity != nil {
		price.Level1Quantity = ConvertInt32ToInt64(input.Level1Quantity)
	}
	if input.Level1Discount != nil {
		price.Level1Discount = ConvertInt32ToInt64(input.Level1Discount)
	}
	if input.Level2Quantity != nil {
		price.Level2Quantity = ConvertInt32ToInt64(input.Level2Quantity)
	}
	if input.Level2Discount != nil {
		price.Level2Discount = ConvertInt32ToInt64(input.Level2Discount)
	}
	if input.Level3Quantity != nil {
		price.Level3Quantity = ConvertInt32ToInt64(input.Level3Quantity)
	}
	if input.Level3Discount != nil {
		price.Level3Discount = ConvertInt32ToInt64(input.Level3Discount)
	}
	if input.Level4Quantity != nil {
		price.Level4Quantity = ConvertInt32ToInt64(input.Level4Quantity)
	}
	if input.Level4Discount != nil {
		price.Level4Discount = ConvertInt32ToInt64(input.Level4Discount)
	}
	if input.PriceHasVat != nil {
		price.PriceHasVAT = ConvertInt32ToInt64(input.PriceHasVat)
	}
	return price
}
func ConvertInt64ToInt32(val *int64) *int32 {
	if val == nil {
		return nil
	}
	t := int32(*val)
	return &t
}
func ConvertInt32ToInt64(val *int32) *int64 {
	if val == nil {
		return nil
	}
	t := int64(*val)
	return &t
}
