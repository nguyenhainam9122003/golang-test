package utils

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"test/graphql_service/model"
	"time"
)

func ConvertToProductModels(input model.CreateProductInput) model.ProductClient {
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
	return model.ProductClient{
		Type:              input.Type,
		UomID:             ConvertStringToUint64(input.UomID),
		IsActive:          isActive,
		SellingStatus:     ConvertSellingStatus(input.SellingStatus),
		Name:              input.Name,
		Code:              input.Code,
		ShortCode:         input.ShortCode,
		ProductType:       ConvertProductTypeToDomain(&input.ProductType),
		VAT:               vat,
		VatType:           ConvertVatType(input.VatType),
		Warranty:          warranty,
		WarrantyUnit:      ConvertWarrantyUnit(input.WarrantyUnit),
		VatPercent:        vatPercent,
		VatValue:          vatValue,
		CategoryID:        ConvertStringToUint64(input.CategoryID),
		ProductGroupID:    ConvertStringToUint64(input.ProductGroupID),
		PlatformID:        platformID,
		CanPreOrder:       canPreOrder,
		CopyrightTerm:     copyrightTerm,
		CopyrightUnit:     ConvertCopyrightUnit(input.CopyrightUnit),
		ImageURL:          input.ImageURL,
		Note:              input.Note,
		MinOrderQuantity:  minOrderQuantity,
		MedusaID:          input.MedusaID,
		OriginalCode:      input.OriginalCode,
		OriginalProductID: ConvertStringToUint64(input.OriginalProductID),
		ProductPrice:      ConvertToProductPricemodels(input.ProductPrice),
	}
}
func ConvertCopyrightUnit(copyrightUnit *model.CopyrightUnit) *model.CopyrightUnit {
	if copyrightUnit == nil {
		return nil
	}
	switch *copyrightUnit {
	case model.CopyrightUnitYear:
		val := model.CopyrightUnitYear
		return &val
	case model.CopyrightUnitMonth:
		val := model.CopyrightUnitMonth
		return &val
	case model.CopyrightUnitDay:
		val := model.CopyrightUnitDay
		return &val
	case model.CopyrightUnitForever:
		val := model.CopyrightUnitForever
		return &val
	default:
		return nil
	}
}
func ConvertWarrantyUnit(warrantyUnit *model.WarrantyUnit) *model.WarrantyUnit {
	if warrantyUnit == nil {
		return nil
	}
	switch *warrantyUnit {
	case model.WarrantyUnitYear:
		val := model.WarrantyUnitYear
		return &val
	case model.WarrantyUnitMonth:
		val := model.WarrantyUnitMonth
		return &val
	case model.WarrantyUnitDay:
		val := model.WarrantyUnitDay
		return &val
	default:
		return nil
	}
}
func ConvertVatType(vatType *model.VatType) *model.VatType {
	if vatType == nil {
		return nil
	}
	switch *vatType {
	case model.VatTypeRequired:
		val := model.VatTypeRequired
		return &val
	case model.VatTypeNotRequired:
		val := model.VatTypeNotRequired
		return &val
	case model.VatTypeOptional:
		val := model.VatTypeOptional
		return &val
	default:
		return nil
	}
}
func ConvertSellingStatus(status *model.SellingStatus) model.SellingStatus {
	if status == nil {
		return model.SellingStatusAvailable
	}
	switch *status {
	case model.SellingStatusAvailable:
		return model.SellingStatusAvailable
	case model.SellingStatusContact:
		return model.SellingStatusContact
	case model.SellingStatusUnavailable:
		return model.SellingStatusUnavailable
	default:
		return model.SellingStatusAvailable
	}
}
func ConvertProductTypeToDomain(productType *model.ProductType) model.ProductType {
	if productType == nil {
		return model.ProductTypeMaterial
	}
	switch *productType {
	case model.ProductTypeMaterial:
		return model.ProductTypeMaterial
	case model.ProductTypeService:
		return model.ProductTypeService
	case model.ProductTypeVoucher:
		return model.ProductTypeVoucher
	case model.ProductTypeKeyLicense:
		return model.ProductTypeKeyLicense
	case model.ProductTypeAccount:
		return model.ProductTypeAccount
	default:
		return model.ProductTypeMaterial
	}
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
func ConvertToProductPricemodels(input *model.CreateProductPriceInput) *model.ProductPriceClient {
	if input == nil {
		return nil
	}
	return &model.ProductPriceClient{
		Level1Price:    int64(input.Level1Price),
		Level2Price:    int64(input.Level2Price),
		Level3Price:    int64(input.Level3Price),
		Level4Price:    int64(input.Level4Price),
		Level1Quantity: IntPtrToInt64Ptr(input.Level1Quantity),
		Level1Discount: IntPtrToInt64Ptr(input.Level1Discount),
		Level2Quantity: IntPtrToInt64Ptr(input.Level2Quantity),
		Level2Discount: IntPtrToInt64Ptr(input.Level2Discount),
		Level3Quantity: IntPtrToInt64Ptr(input.Level3Quantity),
		Level3Discount: IntPtrToInt64Ptr(input.Level3Discount),
		Level4Quantity: IntPtrToInt64Ptr(input.Level4Quantity),
		Level4Discount: IntPtrToInt64Ptr(input.Level4Discount),
		PriceHasVAT:    IntPtrToInt64Ptr(input.PriceHasVat),
	}
}
func IntPtrToInt64Ptr(i *int) *int64 {
	if i == nil {
		return nil
	}
	v := int64(*i)
	return &v
}
func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch val := v.(type) {
	case string:
		return val
	case float64:
		return fmt.Sprintf("%.0f", val)
	case int:
		return strconv.Itoa(val)
	case uint64:
		return strconv.FormatUint(val, 10)
	case fmt.Stringer:
		return val.String()
	default:
		return fmt.Sprintf("%v", val)
	}
}
func ToStringPtr(v interface{}) *string {
	if v == nil {
		return nil
	}
	s := ToString(v) // dùng hàm ToString bạn đã có
	return &s
}
func ConvertProductType(input interface{}) model.ProductType {
	str, ok := input.(string)
	if !ok {
		return model.ProductTypeMaterial
	}
	productType := model.ProductType(str)
	switch productType {
	case model.ProductTypeMaterial:
		return model.ProductTypeMaterial
	case model.ProductTypeService:
		return model.ProductTypeService
	case model.ProductTypeVoucher:
		return model.ProductTypeVoucher
	case model.ProductTypeKeyLicense:
		return model.ProductTypeKeyLicense
	case model.ProductTypeAccount:
		return model.ProductTypeAccount
	default:
		return model.ProductTypeMaterial
	}
}

var ProductFactories = map[string]func() model.Product{
	"MATERIAL":    func() model.Product { return &model.MaterialProduct{} },
	"KEY_LICENSE": func() model.Product { return &model.DigitalProduct{} },
}

// Hàm chuyển đổi snake_case sang PascalCase
func SnakeToPascalCase(name string) string {
	parts := strings.Split(name, "_")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	return strings.Join(parts, "")
}

// Hàm tiền xử lý map để đổi key sang PascalCase
func ConvertKeysToPascalCase(input map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})
	for key, value := range input {
		newKey := SnakeToPascalCase(key)
		if nestedMap, ok := value.(map[string]interface{}); ok {
			// Xử lý đệ quy cho map lồng nhau
			output[newKey] = ConvertKeysToPascalCase(nestedMap)
		} else {
			output[newKey] = value
		}
	}
	return output
}

// Hàm hook để xử lý time.Time
func TimeHook(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if t != reflect.TypeOf(time.Time{}) {
		return data, nil
	}
	if str, ok := data.(string); ok {
		// Hỗ trợ nhiều định dạng thời gian
		for _, layout := range []string{
			time.RFC3339,          // "2006-01-02T15:04:05Z07:00"
			"2006-01-02 15:04:05", // Định dạng tùy chỉnh
			"2006-01-02T15:04:05", // Không có múi giờ
			"2006-01-02",          // Chỉ ngày
		} {
			if t, err := time.Parse(layout, str); err == nil {
				return t, nil
			}
		}
		log.Printf("Failed to parse time string: %v", str)
	}
	return data, nil
}
