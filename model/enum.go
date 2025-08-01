package model

type SellingStatus string
type ProductType string
type VatType string
type WarrantyUnit string
type CopyrightUnit string

const (
	// SellingStatus
	SellingStatusAvailable   SellingStatus = "AVAILABLE"
	SellingStatusContact     SellingStatus = "CONTACT"
	SellingStatusUnavailable SellingStatus = "UNAVAILABLE"

	// ProductType
	ProductTypeMaterial   ProductType = "MATERIAL"
	ProductTypeService    ProductType = "SERVICE"
	ProductTypeVoucher    ProductType = "VOUCHER"
	ProductTypeKeyLicense ProductType = "KEY_LICENSE"
	ProductTypeAccount    ProductType = "ACCOUNT"

	// VAT Type
	VatTypeRequired    VatType = "REQUIRED"
	VatTypeNotRequired VatType = "NOT_REQUIRED"
	VatTypeOptional    VatType = "OPTIONAL"

	// Warranty Unit
	WarrantyUnitYear  WarrantyUnit = "YEAR"
	WarrantyUnitMonth WarrantyUnit = "MONTH"
	WarrantyUnitDay   WarrantyUnit = "DAY"

	// Copyright Unit
	CopyrightUnitYear    CopyrightUnit = "YEAR"
	CopyrightUnitMonth   CopyrightUnit = "MONTH"
	CopyrightUnitDay     CopyrightUnit = "DAY"
	CopyrightUnitForever CopyrightUnit = "FOREVER"
)
