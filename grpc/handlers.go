package main

import (
	"context"
	"log"
	"test/domain"
	"test/model"
	grpcproduct "test/proto/product"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type productServer struct {
	grpcproduct.UnimplementedProductServiceServer
	productService domain.ProductService
}

func (s *productServer) CreateProduct(ctx context.Context, req *grpcproduct.CreateProductRequest) (*grpcproduct.CreateProductResponse, error) {
	log.Printf("📝 Creating product: %s", req.Product.Name)
	
	// Convert proto Product to domain model
	domainProduct := convertProtoToDomainProduct(req.Product)
	
	// Validate product
	if err := domainProduct.Validate(); err != nil {
		log.Printf("❌ Product validation failed: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "product validation failed: %v", err)
	}
	
	// Call service
	err := s.productService.Create(ctx, &domainProduct)
	if err != nil {
		log.Printf("❌ Failed to create product: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create product: %v", err)
	}
	
	log.Printf("✅ Product created successfully with ID: %d", domainProduct.ID)
	
	// Convert back to proto
	protoProduct := convertDomainToProtoProduct(domainProduct)
	return &grpcproduct.CreateProductResponse{
		Product: protoProduct,
	}, nil
}

func (s *productServer) GetProduct(ctx context.Context, req *grpcproduct.GetProductRequest) (*grpcproduct.GetProductResponse, error) {
	log.Printf("🔍 Getting product by ID: %d", req.Id)
	
	// Call service
	domainProduct, err := s.productService.GetByID(ctx, uint(req.Id))
	if err != nil {
		log.Printf("❌ Product not found: %v", err)
		return nil, status.Errorf(codes.NotFound, "product not found: %v", err)
	}
	
	log.Printf("✅ Product found: %s", domainProduct.Name)
	
	// Convert to proto
	protoProduct := convertDomainToProtoProduct(*domainProduct)
	return &grpcproduct.GetProductResponse{
		Product: protoProduct,
	}, nil
}

func (s *productServer) GetProducts(ctx context.Context, req *grpcproduct.GetProductsRequest) (*grpcproduct.GetProductsResponse, error) {
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
	domainProducts, err := s.productService.GetPaginated(ctx, int(req.Limit), offset, query, filter)
	if err != nil {
		log.Printf("❌ Failed to get products: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get products: %v", err)
	}
	
	log.Printf("✅ Found %d products", len(domainProducts))
	
	// Convert to proto
	var protoProducts []*grpcproduct.Product
	for _, p := range domainProducts {
		protoProducts = append(protoProducts, convertDomainToProtoProduct(p))
	}
	
	return &grpcproduct.GetProductsResponse{
		Products: protoProducts,
		Total:    int32(len(domainProducts)),
		Page:     req.Page,
		Limit:    req.Limit,
	}, nil
}

func (s *productServer) UpdateProduct(ctx context.Context, req *grpcproduct.UpdateProductRequest) (*grpcproduct.UpdateProductResponse, error) {
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
	err := s.productService.Update(ctx, req.Id, &domainProduct)
	if err != nil {
		log.Printf("❌ Failed to update product: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to update product: %v", err)
	}
	
	log.Printf("✅ Product updated successfully")
	
	// Convert back to proto
	protoProduct := convertDomainToProtoProduct(domainProduct)
	return &grpcproduct.UpdateProductResponse{
		Product: protoProduct,
	}, nil
}

func (s *productServer) DeleteProduct(ctx context.Context, req *grpcproduct.DeleteProductRequest) (*grpcproduct.DeleteProductResponse, error) {
	log.Printf("🗑️ Deleting product ID: %d", req.Id)
	
	// For now, just return success (implement actual deletion logic if needed)
	log.Printf("✅ Product deletion requested")
	
	return &grpcproduct.DeleteProductResponse{
		Success: true,
	}, nil
} 