package server

import (
	"test/core/model"
	protoProduct "test/proto/gen/product"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Conversion functions
func convertProtoToDomainProduct(protoProduct *protoProduct.Product) model.Product {
	domainProduct := model.Product{
		ID:               protoProduct.Id,
		Type:             protoProduct.Type,
		IsActive:         protoProduct.IsActive,
		Name:             protoProduct.Name,
		Code:             protoProduct.Code,
		VAT:              protoProduct.Vat,
		VatPercent:       int(protoProduct.VatPercent),
		VatValue:         protoProduct.VatValue,
		PlatformID:       protoProduct.PlatformId,
		CanPreOrder:      protoProduct.CanPreOrder,
		MinOrderQuantity: protoProduct.MinOrderQuantity,
	}

	// Handle optional fields
	if protoProduct.UomId != nil {
		domainProduct.UomID = protoProduct.UomId
	}
	if protoProduct.ShortCode != nil {
		domainProduct.ShortCode = protoProduct.ShortCode
	}
	if protoProduct.CategoryId != nil {
		domainProduct.CategoryID = protoProduct.CategoryId
	}
	if protoProduct.ProductGroupId != nil {
		domainProduct.ProductGroupID = protoProduct.ProductGroupId
	}
	if protoProduct.Warranty != nil {
		domainProduct.Warranty = protoProduct.Warranty
	}
	if protoProduct.CopyrightTerm != nil {
		domainProduct.CopyrightTerm = protoProduct.CopyrightTerm
	}
	if protoProduct.ImageUrl != nil {
		domainProduct.ImageURL = protoProduct.ImageUrl
	}
	if protoProduct.Note != nil {
		domainProduct.Note = protoProduct.Note
	}
	if protoProduct.MedusaId != nil {
		domainProduct.MedusaID = protoProduct.MedusaId
	}
	if protoProduct.OriginalCode != nil {
		domainProduct.OriginalCode = protoProduct.OriginalCode
	}
	if protoProduct.OriginalProductId != nil {
		domainProduct.OriginalProductID = protoProduct.OriginalProductId
	}
	if protoProduct.SearchTextV2 != nil {
		domainProduct.SearchTextV2 = protoProduct.SearchTextV2
	}
	if protoProduct.SearchText != nil {
		domainProduct.SearchText = protoProduct.SearchText
	}

	// Convert enums
	domainProduct.SellingStatus = convertProtoSellingStatus(protoProduct.SellingStatus)
	domainProduct.ProductType = convertProtoProductType(protoProduct.ProductType)
	if protoProduct.VatType != nil {
		domainProduct.VatType = convertProtoVatType(*protoProduct.VatType)
	}
	if protoProduct.WarrantyUnit != nil {
		domainProduct.WarrantyUnit = convertProtoWarrantyUnit(*protoProduct.WarrantyUnit)
	}
	if protoProduct.CopyrightUnit != nil {
		domainProduct.CopyrightUnit = convertProtoCopyrightUnit(*protoProduct.CopyrightUnit)
	}

	// Convert timestamps
	if protoProduct.CreatedAt != nil {
		domainProduct.CreatedAt = protoProduct.CreatedAt.AsTime()
	}
	if protoProduct.UpdatedAt != nil {
		domainProduct.UpdatedAt = protoProduct.UpdatedAt.AsTime()
	}

	// Convert ProductPrice
	if protoProduct.ProductPrice != nil {
		domainProduct.ProductPrice = convertProtoToDomainProductPrice(protoProduct.ProductPrice)
	}

	return domainProduct
}

func convertDomainToProtoProduct(domainProduct model.Product) *protoProduct.Product {
	protoProduct := &protoProduct.Product{
		Id:               domainProduct.ID,
		Type:             domainProduct.Type,
		IsActive:         domainProduct.IsActive,
		Name:             domainProduct.Name,
		Code:             domainProduct.Code,
		Vat:              domainProduct.VAT,
		VatPercent:       int32(domainProduct.VatPercent),
		VatValue:         domainProduct.VatValue,
		PlatformId:       domainProduct.PlatformID,
		CanPreOrder:      domainProduct.CanPreOrder,
		MinOrderQuantity: domainProduct.MinOrderQuantity,
	}

	// Handle optional fields
	if domainProduct.UomID != nil {
		protoProduct.UomId = domainProduct.UomID
	}
	if domainProduct.ShortCode != nil {
		protoProduct.ShortCode = domainProduct.ShortCode
	}
	if domainProduct.CategoryID != nil {
		protoProduct.CategoryId = domainProduct.CategoryID
	}
	if domainProduct.ProductGroupID != nil {
		protoProduct.ProductGroupId = domainProduct.ProductGroupID
	}
	if domainProduct.Warranty != nil {
		protoProduct.Warranty = domainProduct.Warranty
	}
	if domainProduct.CopyrightTerm != nil {
		protoProduct.CopyrightTerm = domainProduct.CopyrightTerm
	}
	if domainProduct.ImageURL != nil {
		protoProduct.ImageUrl = domainProduct.ImageURL
	}
	if domainProduct.Note != nil {
		protoProduct.Note = domainProduct.Note
	}
	if domainProduct.MedusaID != nil {
		protoProduct.MedusaId = domainProduct.MedusaID
	}
	if domainProduct.OriginalCode != nil {
		protoProduct.OriginalCode = domainProduct.OriginalCode
	}
	if domainProduct.OriginalProductID != nil {
		protoProduct.OriginalProductId = domainProduct.OriginalProductID
	}
	if domainProduct.SearchTextV2 != nil {
		protoProduct.SearchTextV2 = domainProduct.SearchTextV2
	}
	if domainProduct.SearchText != nil {
		protoProduct.SearchText = domainProduct.SearchText
	}

	// Convert enums
	protoProduct.SellingStatus = convertDomainSellingStatus(domainProduct.SellingStatus)
	protoProduct.ProductType = convertDomainProductType(domainProduct.ProductType)
	if domainProduct.VatType != nil {
		vatType := convertDomainVatType(*domainProduct.VatType)
		protoProduct.VatType = &vatType
	}
	if domainProduct.WarrantyUnit != nil {
		warrantyUnit := convertDomainWarrantyUnit(*domainProduct.WarrantyUnit)
		protoProduct.WarrantyUnit = &warrantyUnit
	}
	if domainProduct.CopyrightUnit != nil {
		copyrightUnit := convertDomainCopyrightUnit(*domainProduct.CopyrightUnit)
		protoProduct.CopyrightUnit = &copyrightUnit
	}

	// Convert timestamps
	protoProduct.CreatedAt = timestamppb.New(domainProduct.CreatedAt)
	protoProduct.UpdatedAt = timestamppb.New(domainProduct.UpdatedAt)

	// Convert ProductPrice
	if domainProduct.ProductPrice != nil {
		protoProduct.ProductPrice = convertDomainToProtoProductPrice(domainProduct.ProductPrice)
	}

	return protoProduct
}

// Enum conversion functions
func convertProtoSellingStatus(status protoProduct.SellingStatus) model.SellingStatus {
	switch status {
	case protoProduct.SellingStatus_SELLING_STATUS_AVAILABLE:
		return model.SellingStatusAvailable
	case protoProduct.SellingStatus_SELLING_STATUS_CONTACT:
		return model.SellingStatusContact
	case protoProduct.SellingStatus_SELLING_STATUS_UNAVAILABLE:
		return model.SellingStatusUnavailable
	default:
		return model.SellingStatusAvailable
	}
}

func convertDomainSellingStatus(status model.SellingStatus) protoProduct.SellingStatus {
	switch status {
	case model.SellingStatusAvailable:
		return protoProduct.SellingStatus_SELLING_STATUS_AVAILABLE
	case model.SellingStatusContact:
		return protoProduct.SellingStatus_SELLING_STATUS_CONTACT
	case model.SellingStatusUnavailable:
		return protoProduct.SellingStatus_SELLING_STATUS_UNAVAILABLE
	default:
		return protoProduct.SellingStatus_SELLING_STATUS_AVAILABLE
	}
}

func convertProtoProductType(productType protoProduct.ProductType) model.ProductType {
	switch productType {
	case protoProduct.ProductType_PRODUCT_TYPE_MATERIAL:
		return model.ProductTypeMaterial
	case protoProduct.ProductType_PRODUCT_TYPE_SERVICE:
		return model.ProductTypeService
	case protoProduct.ProductType_PRODUCT_TYPE_VOUCHER:
		return model.ProductTypeVoucher
	case protoProduct.ProductType_PRODUCT_TYPE_KEY_LICENSE:
		return model.ProductTypeKeyLicense
	case protoProduct.ProductType_PRODUCT_TYPE_ACCOUNT:
		return model.ProductTypeAccount
	default:
		return model.ProductTypeMaterial
	}
}

func convertDomainProductType(productType model.ProductType) protoProduct.ProductType {
	switch productType {
	case model.ProductTypeMaterial:
		return protoProduct.ProductType_PRODUCT_TYPE_MATERIAL
	case model.ProductTypeService:
		return protoProduct.ProductType_PRODUCT_TYPE_SERVICE
	case model.ProductTypeVoucher:
		return protoProduct.ProductType_PRODUCT_TYPE_VOUCHER
	case model.ProductTypeKeyLicense:
		return protoProduct.ProductType_PRODUCT_TYPE_KEY_LICENSE
	case model.ProductTypeAccount:
		return protoProduct.ProductType_PRODUCT_TYPE_ACCOUNT
	default:
		return protoProduct.ProductType_PRODUCT_TYPE_MATERIAL
	}
}

func convertProtoVatType(vatType protoProduct.VatType) *model.VatType {
	switch vatType {
	case protoProduct.VatType_VAT_TYPE_REQUIRED:
		val := model.VatTypeRequired
		return &val
	case protoProduct.VatType_VAT_TYPE_NOT_REQUIRED:
		val := model.VatTypeNotRequired
		return &val
	case protoProduct.VatType_VAT_TYPE_OPTIONAL:
		val := model.VatTypeOptional
		return &val
	default:
		return nil
	}
}

func convertDomainVatType(vatType model.VatType) protoProduct.VatType {
	switch vatType {
	case model.VatTypeRequired:
		return protoProduct.VatType_VAT_TYPE_REQUIRED
	case model.VatTypeNotRequired:
		return protoProduct.VatType_VAT_TYPE_NOT_REQUIRED
	case model.VatTypeOptional:
		return protoProduct.VatType_VAT_TYPE_OPTIONAL
	default:
		return protoProduct.VatType_VAT_TYPE_UNSPECIFIED
	}
}

func convertProtoWarrantyUnit(warrantyUnit protoProduct.WarrantyUnit) *model.WarrantyUnit {
	switch warrantyUnit {
	case protoProduct.WarrantyUnit_WARRANTY_UNIT_YEAR:
		val := model.WarrantyUnitYear
		return &val
	case protoProduct.WarrantyUnit_WARRANTY_UNIT_MONTH:
		val := model.WarrantyUnitMonth
		return &val
	case protoProduct.WarrantyUnit_WARRANTY_UNIT_DAY:
		val := model.WarrantyUnitDay
		return &val
	default:
		return nil
	}
}

func convertDomainWarrantyUnit(warrantyUnit model.WarrantyUnit) protoProduct.WarrantyUnit {
	switch warrantyUnit {
	case model.WarrantyUnitYear:
		return protoProduct.WarrantyUnit_WARRANTY_UNIT_YEAR
	case model.WarrantyUnitMonth:
		return protoProduct.WarrantyUnit_WARRANTY_UNIT_MONTH
	case model.WarrantyUnitDay:
		return protoProduct.WarrantyUnit_WARRANTY_UNIT_DAY
	default:
		return protoProduct.WarrantyUnit_WARRANTY_UNIT_UNSPECIFIED
	}
}

func convertProtoCopyrightUnit(copyrightUnit protoProduct.CopyrightUnit) *model.CopyrightUnit {
	switch copyrightUnit {
	case protoProduct.CopyrightUnit_COPYRIGHT_UNIT_YEAR:
		val := model.CopyrightUnitYear
		return &val
	case protoProduct.CopyrightUnit_COPYRIGHT_UNIT_MONTH:
		val := model.CopyrightUnitMonth
		return &val
	case protoProduct.CopyrightUnit_COPYRIGHT_UNIT_DAY:
		val := model.CopyrightUnitDay
		return &val
	case protoProduct.CopyrightUnit_COPYRIGHT_UNIT_FOREVER:
		val := model.CopyrightUnitForever
		return &val
	default:
		return nil
	}
}

func convertDomainCopyrightUnit(copyrightUnit model.CopyrightUnit) protoProduct.CopyrightUnit {
	switch copyrightUnit {
	case model.CopyrightUnitYear:
		return protoProduct.CopyrightUnit_COPYRIGHT_UNIT_YEAR
	case model.CopyrightUnitMonth:
		return protoProduct.CopyrightUnit_COPYRIGHT_UNIT_MONTH
	case model.CopyrightUnitDay:
		return protoProduct.CopyrightUnit_COPYRIGHT_UNIT_DAY
	case model.CopyrightUnitForever:
		return protoProduct.CopyrightUnit_COPYRIGHT_UNIT_FOREVER
	default:
		return protoProduct.CopyrightUnit_COPYRIGHT_UNIT_UNSPECIFIED
	}
}

func convertProtoToDomainFilter(protoFilter *protoProduct.ProductFilter) model.ProductFilter {
	filter := model.ProductFilter{}

	if protoFilter.ProductType != nil {
		productTypeStr := string(convertProtoProductType(*protoFilter.ProductType))
		filter.ProductType = &productTypeStr
	}
	if protoFilter.SellingStatus != nil {
		sellingStatusStr := string(convertProtoSellingStatus(*protoFilter.SellingStatus))
		filter.SellingStatus = &sellingStatusStr
	}

	return filter
}

func convertProtoToDomainProductPrice(protoPrice *protoProduct.ProductPrice) *model.ProductPrice {
	if protoPrice == nil {
		return nil
	}

	domainPrice := &model.ProductPrice{
		ID:          protoPrice.Id,
		PlatformID:  protoPrice.PlatformId,
		Level1Price: protoPrice.Level1Price,
		Level2Price: protoPrice.Level2Price,
		Level3Price: protoPrice.Level3Price,
		Level4Price: protoPrice.Level4Price,
	}

	// Handle optional fields
	if protoPrice.ProductId != nil {
		domainPrice.ProductID = protoPrice.ProductId
	}
	if protoPrice.Level1Quantity != nil {
		domainPrice.Level1Quantity = protoPrice.Level1Quantity
	}
	if protoPrice.Level1Discount != nil {
		domainPrice.Level1Discount = protoPrice.Level1Discount
	}
	if protoPrice.Level2Quantity != nil {
		domainPrice.Level2Quantity = protoPrice.Level2Quantity
	}
	if protoPrice.Level2Discount != nil {
		domainPrice.Level2Discount = protoPrice.Level2Discount
	}
	if protoPrice.Level3Quantity != nil {
		domainPrice.Level3Quantity = protoPrice.Level3Quantity
	}
	if protoPrice.Level3Discount != nil {
		domainPrice.Level3Discount = protoPrice.Level3Discount
	}
	if protoPrice.Level4Quantity != nil {
		domainPrice.Level4Quantity = protoPrice.Level4Quantity
	}
	if protoPrice.Level4Discount != nil {
		domainPrice.Level4Discount = protoPrice.Level4Discount
	}
	if protoPrice.PriceHasVat != nil {
		domainPrice.PriceHasVAT = protoPrice.PriceHasVat
	}

	// Convert timestamps
	if protoPrice.CreatedAt != nil {
		domainPrice.CreatedAt = protoPrice.CreatedAt.AsTime()
	}
	if protoPrice.UpdatedAt != nil {
		domainPrice.UpdatedAt = protoPrice.UpdatedAt.AsTime()
	}

	return domainPrice
}

func convertDomainToProtoProductPrice(domainPrice *model.ProductPrice) *protoProduct.ProductPrice {
	if domainPrice == nil {
		return nil
	}

	protoPrice := &protoProduct.ProductPrice{
		Id:          domainPrice.ID,
		PlatformId:  domainPrice.PlatformID,
		Level1Price: domainPrice.Level1Price,
		Level2Price: domainPrice.Level2Price,
		Level3Price: domainPrice.Level3Price,
		Level4Price: domainPrice.Level4Price,
	}

	// Handle optional fields
	if domainPrice.ProductID != nil {
		protoPrice.ProductId = domainPrice.ProductID
	}
	if domainPrice.Level1Quantity != nil {
		protoPrice.Level1Quantity = domainPrice.Level1Quantity
	}
	if domainPrice.Level1Discount != nil {
		protoPrice.Level1Discount = domainPrice.Level1Discount
	}
	if domainPrice.Level2Quantity != nil {
		protoPrice.Level2Quantity = domainPrice.Level2Quantity
	}
	if domainPrice.Level2Discount != nil {
		protoPrice.Level2Discount = domainPrice.Level2Discount
	}
	if domainPrice.Level3Quantity != nil {
		protoPrice.Level3Quantity = domainPrice.Level3Quantity
	}
	if domainPrice.Level3Discount != nil {
		protoPrice.Level3Discount = domainPrice.Level3Discount
	}
	if domainPrice.Level4Quantity != nil {
		protoPrice.Level4Quantity = domainPrice.Level4Quantity
	}
	if domainPrice.Level4Discount != nil {
		protoPrice.Level4Discount = domainPrice.Level4Discount
	}
	if domainPrice.PriceHasVAT != nil {
		protoPrice.PriceHasVat = domainPrice.PriceHasVAT
	}

	// Convert timestamps
	protoPrice.CreatedAt = timestamppb.New(domainPrice.CreatedAt)
	protoPrice.UpdatedAt = timestamppb.New(domainPrice.UpdatedAt)

	return protoPrice
}
