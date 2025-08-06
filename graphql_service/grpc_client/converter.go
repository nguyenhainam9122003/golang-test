package grpc_client

import (
	"fmt"
	"strconv"
	"test/graphql_service/model"
	pb "test/proto/gen/product"
)

// Convert GraphQL enums to protobuf enums
func convertGraphQLProductType(productType model.ProductType) pb.ProductType {
	switch productType {
	case model.ProductTypeMaterial:
		return pb.ProductType_PRODUCT_TYPE_MATERIAL
	case model.ProductTypeService:
		return pb.ProductType_PRODUCT_TYPE_SERVICE
	case model.ProductTypeVoucher:
		return pb.ProductType_PRODUCT_TYPE_VOUCHER
	case model.ProductTypeKeyLicense:
		return pb.ProductType_PRODUCT_TYPE_KEY_LICENSE
	case model.ProductTypeAccount:
		return pb.ProductType_PRODUCT_TYPE_ACCOUNT
	default:
		return pb.ProductType_PRODUCT_TYPE_MATERIAL
	}
}

func convertGraphQLSellingStatus(sellingStatus model.SellingStatus) pb.SellingStatus {
	switch sellingStatus {
	case model.SellingStatusAvailable:
		return pb.SellingStatus_SELLING_STATUS_AVAILABLE
	case model.SellingStatusContact:
		return pb.SellingStatus_SELLING_STATUS_CONTACT
	case model.SellingStatusUnavailable:
		return pb.SellingStatus_SELLING_STATUS_UNAVAILABLE
	default:
		return pb.SellingStatus_SELLING_STATUS_AVAILABLE
	}
}

func convertGraphQLVatType(vatType model.VatType) pb.VatType {
	switch vatType {
	case model.VatTypeRequired:
		return pb.VatType_VAT_TYPE_REQUIRED
	case model.VatTypeNotRequired:
		return pb.VatType_VAT_TYPE_NOT_REQUIRED
	case model.VatTypeOptional:
		return pb.VatType_VAT_TYPE_OPTIONAL
	default:
		return pb.VatType_VAT_TYPE_UNSPECIFIED
	}
}

func convertGraphQLWarrantyUnit(warrantyUnit model.WarrantyUnit) pb.WarrantyUnit {
	switch warrantyUnit {
	case model.WarrantyUnitYear:
		return pb.WarrantyUnit_WARRANTY_UNIT_YEAR
	case model.WarrantyUnitMonth:
		return pb.WarrantyUnit_WARRANTY_UNIT_MONTH
	case model.WarrantyUnitDay:
		return pb.WarrantyUnit_WARRANTY_UNIT_DAY
	default:
		return pb.WarrantyUnit_WARRANTY_UNIT_UNSPECIFIED
	}
}

func convertGraphQLCopyrightUnit(copyrightUnit model.CopyrightUnit) pb.CopyrightUnit {
	switch copyrightUnit {
	case model.CopyrightUnitYear:
		return pb.CopyrightUnit_COPYRIGHT_UNIT_YEAR
	case model.CopyrightUnitMonth:
		return pb.CopyrightUnit_COPYRIGHT_UNIT_MONTH
	case model.CopyrightUnitDay:
		return pb.CopyrightUnit_COPYRIGHT_UNIT_DAY
	case model.CopyrightUnitForever:
		return pb.CopyrightUnit_COPYRIGHT_UNIT_FOREVER
	default:
		return pb.CopyrightUnit_COPYRIGHT_UNIT_UNSPECIFIED
	}
}

// Convert protobuf enums to GraphQL enums
func convertProtoProductType(productType pb.ProductType) model.ProductType {
	switch productType {
	case pb.ProductType_PRODUCT_TYPE_MATERIAL:
		return model.ProductTypeMaterial
	case pb.ProductType_PRODUCT_TYPE_SERVICE:
		return model.ProductTypeService
	case pb.ProductType_PRODUCT_TYPE_VOUCHER:
		return model.ProductTypeVoucher
	case pb.ProductType_PRODUCT_TYPE_KEY_LICENSE:
		return model.ProductTypeKeyLicense
	case pb.ProductType_PRODUCT_TYPE_ACCOUNT:
		return model.ProductTypeAccount
	default:
		return model.ProductTypeMaterial
	}
}

func convertProtoSellingStatus(sellingStatus pb.SellingStatus) model.SellingStatus {
	switch sellingStatus {
	case pb.SellingStatus_SELLING_STATUS_AVAILABLE:
		return model.SellingStatusAvailable
	case pb.SellingStatus_SELLING_STATUS_CONTACT:
		return model.SellingStatusContact
	case pb.SellingStatus_SELLING_STATUS_UNAVAILABLE:
		return model.SellingStatusUnavailable
	default:
		return model.SellingStatusAvailable
	}
}

func convertProtoVatType(vatType pb.VatType) model.VatType {
	switch vatType {
	case pb.VatType_VAT_TYPE_REQUIRED:
		return model.VatTypeRequired
	case pb.VatType_VAT_TYPE_NOT_REQUIRED:
		return model.VatTypeNotRequired
	case pb.VatType_VAT_TYPE_OPTIONAL:
		return model.VatTypeOptional
	default:
		return model.VatTypeNotRequired
	}
}

func convertProtoWarrantyUnit(warrantyUnit pb.WarrantyUnit) model.WarrantyUnit {
	switch warrantyUnit {
	case pb.WarrantyUnit_WARRANTY_UNIT_YEAR:
		return model.WarrantyUnitYear
	case pb.WarrantyUnit_WARRANTY_UNIT_MONTH:
		return model.WarrantyUnitMonth
	case pb.WarrantyUnit_WARRANTY_UNIT_DAY:
		return model.WarrantyUnitDay
	default:
		return model.WarrantyUnitDay
	}
}

func convertProtoCopyrightUnit(copyrightUnit pb.CopyrightUnit) model.CopyrightUnit {
	switch copyrightUnit {
	case pb.CopyrightUnit_COPYRIGHT_UNIT_YEAR:
		return model.CopyrightUnitYear
	case pb.CopyrightUnit_COPYRIGHT_UNIT_MONTH:
		return model.CopyrightUnitMonth
	case pb.CopyrightUnit_COPYRIGHT_UNIT_DAY:
		return model.CopyrightUnitDay
	case pb.CopyrightUnit_COPYRIGHT_UNIT_FOREVER:
		return model.CopyrightUnitForever
	default:
		return model.CopyrightUnitYear
	}
}

// Convert string ID to uint64
func convertStringToUint64(id *string) (*uint64, error) {
	if id == nil {
		return nil, nil
	}
	parsedID, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}
	return &parsedID, nil
}

// Convert int to int32
func convertIntToInt32(val *int) *int32 {
	if val == nil {
		return nil
	}
	val32 := int32(*val)
	return &val32
}

// Convert int32 to int
func convertInt32ToInt(val *int32) *int {
	if val == nil {
		return nil
	}
	valInt := int(*val)
	return &valInt
}

// Convert ProductPrice from GraphQL to protobuf
func convertGraphQLProductPriceToProto(productPrice *model.CreateProductPriceInput) *pb.ProductPrice {
	if productPrice == nil {
		return nil
	}

	pbProductPrice := &pb.ProductPrice{
		PlatformId:  1, // Default platform ID
		Level1Price: int64(productPrice.Level1Price),
		Level2Price: int64(productPrice.Level2Price),
		Level3Price: int64(productPrice.Level3Price),
		Level4Price: int64(productPrice.Level4Price),
	}

	// Convert optional fields
	if productPrice.Level1Quantity != nil {
		level1Quantity := int64(*productPrice.Level1Quantity)
		pbProductPrice.Level1Quantity = &level1Quantity
	}
	if productPrice.Level1Discount != nil {
		level1Discount := int64(*productPrice.Level1Discount)
		pbProductPrice.Level1Discount = &level1Discount
	}
	if productPrice.Level2Quantity != nil {
		level2Quantity := int64(*productPrice.Level2Quantity)
		pbProductPrice.Level2Quantity = &level2Quantity
	}
	if productPrice.Level2Discount != nil {
		level2Discount := int64(*productPrice.Level2Discount)
		pbProductPrice.Level2Discount = &level2Discount
	}
	if productPrice.Level3Quantity != nil {
		level3Quantity := int64(*productPrice.Level3Quantity)
		pbProductPrice.Level3Quantity = &level3Quantity
	}
	if productPrice.Level3Discount != nil {
		level3Discount := int64(*productPrice.Level3Discount)
		pbProductPrice.Level3Discount = &level3Discount
	}
	if productPrice.Level4Quantity != nil {
		level4Quantity := int64(*productPrice.Level4Quantity)
		pbProductPrice.Level4Quantity = &level4Quantity
	}
	if productPrice.Level4Discount != nil {
		level4Discount := int64(*productPrice.Level4Discount)
		pbProductPrice.Level4Discount = &level4Discount
	}
	if productPrice.PriceHasVat != nil {
		priceHasVat := int64(*productPrice.PriceHasVat)
		pbProductPrice.PriceHasVat = &priceHasVat
	}

	return pbProductPrice
}

// Convert ProductPrice from protobuf to GraphQL
func convertProtoProductPriceToGraphQL(pbPrice *pb.ProductPrice) *model.ProductPrice {
	if pbPrice == nil {
		return nil
	}

	productPrice := &model.ProductPrice{
		ID:             strconv.FormatUint(pbPrice.Id, 10),
		PlatformID:     strconv.FormatUint(pbPrice.PlatformId, 10),
		Level1Price:    int(pbPrice.Level1Price),
		Level2Price:    int(pbPrice.Level2Price),
		Level3Price:    int(pbPrice.Level3Price),
		Level4Price:    int(pbPrice.Level4Price),
		Level1Quantity: nil,
		Level1Discount: nil,
		Level2Quantity: nil,
		Level2Discount: nil,
		Level3Quantity: nil,
		Level3Discount: nil,
		Level4Quantity: nil,
		Level4Discount: nil,
		PriceHasVat:    nil,
		CreatedAt:      pbPrice.CreatedAt.AsTime(),
		UpdatedAt:      pbPrice.UpdatedAt.AsTime(),
	}

	// Set optional fields for ProductPrice
	if pbPrice.Level1Quantity != nil {
		level1Quantity := int(*pbPrice.Level1Quantity)
		productPrice.Level1Quantity = &level1Quantity
	}
	if pbPrice.Level1Discount != nil {
		level1Discount := int(*pbPrice.Level1Discount)
		productPrice.Level1Discount = &level1Discount
	}
	if pbPrice.Level2Quantity != nil {
		level2Quantity := int(*pbPrice.Level2Quantity)
		productPrice.Level2Quantity = &level2Quantity
	}
	if pbPrice.Level2Discount != nil {
		level2Discount := int(*pbPrice.Level2Discount)
		productPrice.Level2Discount = &level2Discount
	}
	if pbPrice.Level3Quantity != nil {
		level3Quantity := int(*pbPrice.Level3Quantity)
		productPrice.Level3Quantity = &level3Quantity
	}
	if pbPrice.Level3Discount != nil {
		level3Discount := int(*pbPrice.Level3Discount)
		productPrice.Level3Discount = &level3Discount
	}
	if pbPrice.Level4Quantity != nil {
		level4Quantity := int(*pbPrice.Level4Quantity)
		productPrice.Level4Quantity = &level4Quantity
	}
	if pbPrice.Level4Discount != nil {
		level4Discount := int(*pbPrice.Level4Discount)
		productPrice.Level4Discount = &level4Discount
	}
	if pbPrice.PriceHasVat != nil {
		priceHasVat := int(*pbPrice.PriceHasVat)
		productPrice.PriceHasVat = &priceHasVat
	}

	return productPrice
}
