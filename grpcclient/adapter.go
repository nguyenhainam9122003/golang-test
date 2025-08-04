package grpcclient

import (
	"context"
	"log"
	"test/domain"
	"test/model"
)

// ProductServiceAdapter implements domain.ProductService interface
type ProductServiceAdapter struct {
	client *ProductClient
}

// NewProductServiceAdapter creates a new ProductServiceAdapter that implements domain.ProductService
func NewProductServiceAdapter(client *ProductClient) domain.ProductService {
	return &ProductServiceAdapter{
		client: client,
	}
}

// Create implements domain.ProductService.Create
func (a *ProductServiceAdapter) Create(ctx context.Context, product *model.Product) error {
	log.Printf("📤 gRPC Client: Creating product: %s", product.Name)
	err := a.client.CreateProduct(ctx, product)
	if err != nil {
		log.Printf("❌ gRPC Client: Failed to create product: %v", err)
		return err
	}
	log.Printf("✅ gRPC Client: Product created successfully")
	return nil
}

// GetAll implements domain.ProductService.GetAll
func (a *ProductServiceAdapter) GetAll(ctx context.Context) ([]model.Product, error) {
	log.Printf("📤 gRPC Client: Getting all products")
	products, err := a.client.GetProducts(ctx, 1, 100, "", nil)
	if err != nil {
		log.Printf("❌ gRPC Client: Failed to get products: %v", err)
		return nil, err
	}
	log.Printf("✅ gRPC Client: Retrieved %d products", len(products))
	return products, nil
}

// GetByID implements domain.ProductService.GetByID
func (a *ProductServiceAdapter) GetByID(ctx context.Context, id uint) (*model.Product, error) {
	log.Printf("📤 gRPC Client: Getting product by ID: %d", id)
	product, err := a.client.GetProduct(ctx, id)
	if err != nil {
		log.Printf("❌ gRPC Client: Failed to get product: %v", err)
		return nil, err
	}
	log.Printf("✅ gRPC Client: Product found: %s", product.Name)
	return product, nil
}

// GetPaginated implements domain.ProductService.GetPaginated
func (a *ProductServiceAdapter) GetPaginated(ctx context.Context, limit, offset int, query string, filter model.ProductFilter) ([]model.Product, error) {
	page := (offset / limit) + 1
	log.Printf("📤 gRPC Client: Getting products - Page: %d, Limit: %d", page, limit)
	products, err := a.client.GetProducts(ctx, int(page), limit, query, &filter)
	if err != nil {
		log.Printf("❌ gRPC Client: Failed to get products: %v", err)
		return nil, err
	}
	log.Printf("✅ gRPC Client: Retrieved %d products", len(products))
	return products, nil
}

// Update implements domain.ProductService.Update
func (a *ProductServiceAdapter) Update(ctx context.Context, id uint64, product *model.Product) error {
	log.Printf("📤 gRPC Client: Updating product ID: %d", id)
	err := a.client.UpdateProduct(ctx, id, product)
	if err != nil {
		log.Printf("❌ gRPC Client: Failed to update product: %v", err)
		return err
	}
	log.Printf("✅ gRPC Client: Product updated successfully")
	return nil
} 