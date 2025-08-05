package server

import (
	"context"
	"log"
	"test/grpc_service/domain"
	"test/grpc_service/model"
	protoProduct "test/proto/gen/product"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductServer struct {
	protoProduct.UnimplementedProductServiceServer
	ProductService domain.ProductService
}

func (s *ProductServer) CreateProduct(ctx context.Context, req *protoProduct.CreateProductRequest) (*protoProduct.CreateProductResponse, error) {
	log.Printf("📝 Creating product: %s", req.Product.Name)

	// Convert proto Product to domain model
	domainProduct := convertProtoToDomainProduct(req.Product)

	// Validate product
	if err := domainProduct.Validate(); err != nil {
		log.Printf("❌ Product validation failed: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "product validation failed: %v", err)
	}

	// Call service
	err := s.ProductService.Create(ctx, &domainProduct)
	if err != nil {
		log.Printf("❌ Failed to create product: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create product: %v", err)
	}

	log.Printf("✅ Product created successfully with ID: %d", domainProduct.ID)

	// Convert back to proto
	myProtoProduct := convertDomainToProtoProduct(domainProduct)
	return &protoProduct.CreateProductResponse{
		Product: myProtoProduct,
	}, nil
}

func (s *ProductServer) GetProduct(ctx context.Context, req *protoProduct.GetProductRequest) (*protoProduct.GetProductResponse, error) {
	log.Printf("🔍 Getting product by ID: %d", req.Id)

	// Call service
	domainProduct, err := s.ProductService.GetByID(ctx, uint(req.Id))
	if err != nil {
		log.Printf("❌ Product not found: %v", err)
		return nil, status.Errorf(codes.NotFound, "product not found: %v", err)
	}

	log.Printf("✅ Product found: %s", domainProduct.Name)

	// Convert to proto
	myProtoProduct := convertDomainToProtoProduct(*domainProduct)
	return &protoProduct.GetProductResponse{
		Product: myProtoProduct,
	}, nil
}

func (s *ProductServer) GetProducts(ctx context.Context, req *protoProduct.GetProductsRequest) (*protoProduct.GetProductsResponse, error) {
	log.Printf("📋 Getting products - Page: %d, Limit: %d", req.Page, req.Limit)

	// Convert filter
	var filter model.ProductFilter
	if req.Filter != nil {
		filter = convertProtoToDomainFilter(req.Filter)
	}

	// Handle optional query
	var query string
	if req.Query != nil {
		query = *req.Query
	}

	// Calculate offset
	offset := (int(req.Page) - 1) * int(req.Limit)

	// Call service
	domainProducts, err := s.ProductService.GetPaginated(ctx, int(req.Limit), offset, query, filter)
	if err != nil {
		log.Printf("❌ Failed to get products: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get products: %v", err)
	}

	log.Printf("✅ Found %d products", len(domainProducts))

	// Convert to proto
	var protoProducts []*protoProduct.Product
	for _, p := range domainProducts {
		protoProducts = append(protoProducts, convertDomainToProtoProduct(p))
	}

	return &protoProduct.GetProductsResponse{
		Products: protoProducts,
		Total:    int32(len(domainProducts)),
		Page:     req.Page,
		Limit:    req.Limit,
	}, nil
}

func (s *ProductServer) UpdateProduct(ctx context.Context, req *protoProduct.UpdateProductRequest) (*protoProduct.UpdateProductResponse, error) {
	log.Printf("🔄 Updating product ID: %d", req.Id)

	// Convert proto Product to domain model
	domainProduct := convertProtoToDomainProduct(req.Product)
	domainProduct.ID = req.Id

	// Validate product
	if err := domainProduct.Validate(); err != nil {
		log.Printf("❌ Product validation failed: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "product validation failed: %v", err)
	}

	// Call service
	err := s.ProductService.Update(ctx, req.Id, &domainProduct)
	if err != nil {
		log.Printf("❌ Failed to update product: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to update product: %v", err)
	}

	log.Printf("✅ Product updated successfully")

	// Convert back to proto
	myProtoProduct := convertDomainToProtoProduct(domainProduct)
	return &protoProduct.UpdateProductResponse{
		Product: myProtoProduct,
	}, nil
}
