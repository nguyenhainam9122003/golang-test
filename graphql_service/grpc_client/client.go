package grpc_client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"test/graphql_service/model"
	pb "test/proto/gen/product"
)

type ProductClient struct {
	client pb.ProductServiceClient
}

func NewProductClient(conn *grpc.ClientConn) *ProductClient {
	return &ProductClient{
		client: pb.NewProductServiceClient(conn),
	}
}

func (pc *ProductClient) GetProduct(ctx context.Context, id string) (model.Product, error) {
	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	resp, err := pc.client.GetProduct(ctx, &pb.GetProductRequest{Id: parsedID})
	if err != nil {
		log.Printf("gRPC error: %v", err)
		return nil, err
	}

	p := resp.GetProduct()
	return pc.convertProduct(p)
}

func (pc *ProductClient) GetProducts(ctx context.Context, page *int32, limit *int32, query *string, filter *model.ProductFilter) (*model.ProductPagination, error) {
	// Set default values
	pageNum := int32(1)
	if page != nil {
		pageNum = *page
	}
	
	limitNum := int32(10)
	if limit != nil {
		limitNum = *limit
	}
	
	// Convert filter
	var pbFilter *pb.ProductFilter
	if filter != nil {
		pbFilter = &pb.ProductFilter{}
		
		if filter.ProductType != nil {
			var productType pb.ProductType
			switch *filter.ProductType {
			case model.ProductTypeMaterial:
				productType = pb.ProductType_PRODUCT_TYPE_MATERIAL
			case model.ProductTypeService:
				productType = pb.ProductType_PRODUCT_TYPE_SERVICE
			case model.ProductTypeVoucher:
				productType = pb.ProductType_PRODUCT_TYPE_VOUCHER
			case model.ProductTypeKeyLicense:
				productType = pb.ProductType_PRODUCT_TYPE_KEY_LICENSE
			case model.ProductTypeAccount:
				productType = pb.ProductType_PRODUCT_TYPE_ACCOUNT
			default:
				productType = pb.ProductType_PRODUCT_TYPE_UNSPECIFIED
			}
			pbFilter.ProductType = &productType
		}
		
		if filter.SellingStatus != nil {
			var sellingStatus pb.SellingStatus
			switch *filter.SellingStatus {
			case model.SellingStatusAvailable:
				sellingStatus = pb.SellingStatus_SELLING_STATUS_AVAILABLE
			case model.SellingStatusContact:
				sellingStatus = pb.SellingStatus_SELLING_STATUS_CONTACT
			case model.SellingStatusUnavailable:
				sellingStatus = pb.SellingStatus_SELLING_STATUS_UNAVAILABLE
			default:
				sellingStatus = pb.SellingStatus_SELLING_STATUS_UNSPECIFIED
			}
			pbFilter.SellingStatus = &sellingStatus
		}
	}
	
	resp, err := pc.client.GetProducts(ctx, &pb.GetProductsRequest{
		Page:   pageNum,
		Limit:  limitNum,
		Query:  query,
		Filter: pbFilter,
	})
	if err != nil {
		log.Printf("gRPC error: %v", err)
		return nil, err
	}
	
	// Convert products
	var products []model.Product
	for _, p := range resp.Products {
		product, err := pc.convertProduct(p)
		if err != nil {
			log.Printf("Error converting product: %v", err)
			continue
		}
		products = append(products, product)
	}
	
	return &model.ProductPagination{
		Page:  int(pageNum),
		Limit: int(limitNum),
		Items: products,
		Total: int(resp.Total),
	}, nil
}

func (pc *ProductClient) CreateProduct(ctx context.Context, input model.CreateProductInput) (model.Product, error) {
	log.Printf("🔧 Creating product with input: %+v", input)
	
	// Convert CreateProductInput to protobuf Product
	pbProduct := &pb.Product{
		Type:     input.Type,
		Name:     input.Name,
		Code:     input.Code,
		PlatformId: 1, // Default platform ID
	}

	// Set default values and handle optional fields
	if input.IsActive != nil {
		pbProduct.IsActive = *input.IsActive
	} else {
		pbProduct.IsActive = true // Default value
	}

	if input.ShortCode != nil {
		pbProduct.ShortCode = input.ShortCode
	}

	if input.Vat != nil {
		pbProduct.Vat = *input.Vat
	} else {
		pbProduct.Vat = false // Default value
	}

	if input.VatPercent != nil {
		pbProduct.VatPercent = int32(*input.VatPercent)
	} else {
		pbProduct.VatPercent = 0 // Default value
	}

	if input.VatValue != nil {
		pbProduct.VatValue = int64(*input.VatValue)
	} else {
		pbProduct.VatValue = 0 // Default value
	}

	if input.CanPreOrder != nil {
		pbProduct.CanPreOrder = *input.CanPreOrder
	} else {
		pbProduct.CanPreOrder = true // Default value
	}

	if input.ImageURL != nil {
		pbProduct.ImageUrl = input.ImageURL
	}

	if input.Note != nil {
		pbProduct.Note = input.Note
	}

	if input.MinOrderQuantity != nil {
		pbProduct.MinOrderQuantity = int64(*input.MinOrderQuantity)
	} else {
		pbProduct.MinOrderQuantity = 1 // Default value
	}

	if input.MedusaID != nil {
		pbProduct.MedusaId = input.MedusaID
	}

	if input.OriginalCode != nil {
		pbProduct.OriginalCode = input.OriginalCode
	}

	// Convert enums
	pbProduct.ProductType = convertGraphQLProductType(input.ProductType)

	if input.SellingStatus != nil {
		pbProduct.SellingStatus = convertGraphQLSellingStatus(*input.SellingStatus)
	} else {
		pbProduct.SellingStatus = pb.SellingStatus_SELLING_STATUS_AVAILABLE // Default value
	}

	// Convert optional fields
	if input.UomID != nil {
		uomID, err := strconv.ParseUint(*input.UomID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid uomId: %v", err)
		}
		pbProduct.UomId = &uomID
	}

	if input.VatType != nil {
		vt := convertGraphQLVatType(*input.VatType)
		pbProduct.VatType = &vt
	}

	if input.Warranty != nil {
		warranty := uint64(*input.Warranty)
		pbProduct.Warranty = &warranty
	}

	if input.WarrantyUnit != nil {
		wu := convertGraphQLWarrantyUnit(*input.WarrantyUnit)
		pbProduct.WarrantyUnit = &wu
	}

	if input.CategoryID != nil {
		categoryID, err := strconv.ParseUint(*input.CategoryID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid categoryId: %v", err)
		}
		pbProduct.CategoryId = &categoryID
	}

	if input.ProductGroupID != nil {
		productGroupID, err := strconv.ParseUint(*input.ProductGroupID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid productGroupId: %v", err)
		}
		pbProduct.ProductGroupId = &productGroupID
	}

	if input.PlatformID != nil {
		platformID, err := strconv.ParseUint(*input.PlatformID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid platformId: %v", err)
		}
		pbProduct.PlatformId = platformID
	}

	if input.CopyrightTerm != nil {
		copyrightTerm := uint64(*input.CopyrightTerm)
		pbProduct.CopyrightTerm = &copyrightTerm
	}

	if input.CopyrightUnit != nil {
		cu := convertGraphQLCopyrightUnit(*input.CopyrightUnit)
		pbProduct.CopyrightUnit = &cu
	}

	if input.OriginalProductID != nil {
		originalProductID, err := strconv.ParseUint(*input.OriginalProductID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid originalProductId: %v", err)
		}
		pbProduct.OriginalProductId = &originalProductID
	}

	// Convert ProductPrice if provided
	if input.ProductPrice != nil {
		pbProduct.ProductPrice = convertGraphQLProductPriceToProto(input.ProductPrice)
	}

	// Call gRPC service
	log.Printf("📞 Calling gRPC CreateProduct with: %+v", pbProduct)
	resp, err := pc.client.CreateProduct(ctx, &pb.CreateProductRequest{
		Product: pbProduct,
	})
	if err != nil {
		log.Printf("❌ gRPC error: %v", err)
		return nil, err
	}

	log.Printf("✅ gRPC CreateProduct success, converting response...")
	// Convert response back to GraphQL model
	return pc.convertProduct(resp.GetProduct())
}

func (pc *ProductClient) convertProduct(p *pb.Product) (model.Product, error) {
	// Convert common fields
	productID := strconv.FormatUint(p.Id, 10)
	platformID := strconv.FormatUint(p.PlatformId, 10)
	
	var uomID *string
	if p.UomId != nil {
		uomIDStr := strconv.FormatUint(*p.UomId, 10)
		uomID = &uomIDStr
	}
	
	var categoryID *string
	if p.CategoryId != nil {
		categoryIDStr := strconv.FormatUint(*p.CategoryId, 10)
		categoryID = &categoryIDStr
	}
	
	var productGroupID *string
	if p.ProductGroupId != nil {
		productGroupIDStr := strconv.FormatUint(*p.ProductGroupId, 10)
		productGroupID = &productGroupIDStr
	}
	
	var warranty *int
	if p.Warranty != nil {
		warrantyInt := int(*p.Warranty)
		warranty = &warrantyInt
	}
	
	var copyrightTerm *int
	if p.CopyrightTerm != nil {
		copyrightTermInt := int(*p.CopyrightTerm)
		copyrightTerm = &copyrightTermInt
	}
	
	var originalProductID *string
	if p.OriginalProductId != nil {
		originalProductIDStr := strconv.FormatUint(*p.OriginalProductId, 10)
		originalProductID = &originalProductIDStr
	}
	
	// Convert ProductPrice if exists
	var productPrice *model.ProductPrice
	if p.ProductPrice != nil {
		productPrice = convertProtoProductPriceToGraphQL(p.ProductPrice)
		productPrice.ProductID = &productID
	}
	
	// Convert enums
	sellingStatus := convertProtoSellingStatus(p.SellingStatus)
	productType := convertProtoProductType(p.ProductType)
	
	var vatType *model.VatType
	if p.VatType != nil {
		vt := convertProtoVatType(*p.VatType)
		vatType = &vt
	}
	
	var warrantyUnit *model.WarrantyUnit
	if p.WarrantyUnit != nil {
		wu := convertProtoWarrantyUnit(*p.WarrantyUnit)
		warrantyUnit = &wu
	}
	
	var copyrightUnit *model.CopyrightUnit
	if p.CopyrightUnit != nil {
		cu := convertProtoCopyrightUnit(*p.CopyrightUnit)
		copyrightUnit = &cu
	}
	
	// Determine product type and return appropriate struct
	switch p.ProductType {
	case pb.ProductType_PRODUCT_TYPE_MATERIAL:
		var shortCode *string
		if p.ShortCode != nil {
			shortCode = p.ShortCode
		}
		
		return &model.MaterialProduct{
			ShortCode:         shortCode,
			ID:                productID,
			CreatedAt:         p.CreatedAt.AsTime(),
			UpdatedAt:         p.UpdatedAt.AsTime(),
			Type:              p.Type,
			UomID:             uomID,
			IsActive:          p.IsActive,
			SellingStatus:     sellingStatus,
			Name:              p.Name,
			Code:              p.Code,
			ProductType:       productType,
			Vat:               p.Vat,
			VatType:           vatType,
			Warranty:          warranty,
			WarrantyUnit:      warrantyUnit,
			VatPercent:        int(p.VatPercent),
			VatValue:          int(p.VatValue),
			CategoryID:        categoryID,
			ProductGroupID:    productGroupID,
			PlatformID:        platformID,
			CanPreOrder:       p.CanPreOrder,
			CopyrightTerm:     copyrightTerm,
			CopyrightUnit:     copyrightUnit,
			ImageURL:          p.ImageUrl,
			Note:              p.Note,
			MinOrderQuantity:  int(p.MinOrderQuantity),
			MedusaID:          p.MedusaId,
			OriginalProductID: originalProductID,
			SearchTextV2:      p.SearchTextV2,
			SearchText:        p.SearchText,
			ProductPrice:      productPrice,
		}, nil
		
	case pb.ProductType_PRODUCT_TYPE_SERVICE, pb.ProductType_PRODUCT_TYPE_VOUCHER, 
		 pb.ProductType_PRODUCT_TYPE_KEY_LICENSE, pb.ProductType_PRODUCT_TYPE_ACCOUNT:
		var originalCode *string
		if p.OriginalCode != nil {
			originalCode = p.OriginalCode
		}
		
		return &model.DigitalProduct{
			OriginalCode:      originalCode,
			ID:                productID,
			CreatedAt:         p.CreatedAt.AsTime(),
			UpdatedAt:         p.UpdatedAt.AsTime(),
			Type:              p.Type,
			UomID:             uomID,
			IsActive:          p.IsActive,
			SellingStatus:     sellingStatus,
			Name:              p.Name,
			Code:              p.Code,
			ProductType:       productType,
			Vat:               p.Vat,
			VatType:           vatType,
			Warranty:          warranty,
			WarrantyUnit:      warrantyUnit,
			VatPercent:        int(p.VatPercent),
			VatValue:          int(p.VatValue),
			CategoryID:        categoryID,
			ProductGroupID:    productGroupID,
			PlatformID:        platformID,
			CanPreOrder:       p.CanPreOrder,
			CopyrightTerm:     copyrightTerm,
			CopyrightUnit:     copyrightUnit,
			ImageURL:          p.ImageUrl,
			Note:              p.Note,
			MinOrderQuantity:  int(p.MinOrderQuantity),
			MedusaID:          p.MedusaId,
			OriginalProductID: originalProductID,
			SearchTextV2:      p.SearchTextV2,
			SearchText:        p.SearchText,
			ProductPrice:      productPrice,
		}, nil
		
	default:
		// Fallback to MaterialProduct for unknown types
		return &model.MaterialProduct{
			ID:                productID,
			CreatedAt:         p.CreatedAt.AsTime(),
			UpdatedAt:         p.UpdatedAt.AsTime(),
			Type:              p.Type,
			UomID:             uomID,
			IsActive:          p.IsActive,
			SellingStatus:     sellingStatus,
			Name:              p.Name,
			Code:              p.Code,
			ProductType:       productType,
			Vat:               p.Vat,
			VatType:           vatType,
			Warranty:          warranty,
			WarrantyUnit:      warrantyUnit,
			VatPercent:        int(p.VatPercent),
			VatValue:          int(p.VatValue),
			CategoryID:        categoryID,
			ProductGroupID:    productGroupID,
			PlatformID:        platformID,
			CanPreOrder:       p.CanPreOrder,
			CopyrightTerm:     copyrightTerm,
			CopyrightUnit:     copyrightUnit,
			ImageURL:          p.ImageUrl,
			Note:              p.Note,
			MinOrderQuantity:  int(p.MinOrderQuantity),
			MedusaID:          p.MedusaId,
			OriginalProductID: originalProductID,
			SearchTextV2:      p.SearchTextV2,
			SearchText:        p.SearchText,
			ProductPrice:      productPrice,
		}, nil
	}
}
