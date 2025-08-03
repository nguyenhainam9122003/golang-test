package http_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"test/graphql_service/models"
	"time"
)

type ProductHTTPClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewProductHTTPClient(baseURL string) *ProductHTTPClient {
	return &ProductHTTPClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetAllProducts gọi API để lấy tất cả products
func (c *ProductHTTPClient) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	url := fmt.Sprintf("%s/products", c.baseURL)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var apiResponse models.APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert API response to our models
	var products []models.Product
	if data, ok := apiResponse.Data.([]interface{}); ok {
		for _, item := range data {
			if productData, ok := item.(map[string]interface{}); ok {
				product := c.convertMapToProduct(productData)
				products = append(products, product)
			}
		}
	}

	return products, nil
}

// GetProductByID gọi API để lấy product theo ID
func (c *ProductHTTPClient) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	url := fmt.Sprintf("%s/products/%s", c.baseURL, id)
	
	fmt.Printf("Calling API: %s\n", url)
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	fmt.Printf("API Response Status: %d\n", resp.StatusCode)

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("product not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var apiResponse models.APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	fmt.Printf("API Response: %+v\n", apiResponse)

	// Convert API response to our model
	if productData, ok := apiResponse.Data.(map[string]interface{}); ok {
		fmt.Printf("Product Data: %+v\n", productData)
		product := c.convertMapToProduct(productData)
		fmt.Printf("Converted Product: %+v\n", product)
		return &product, nil
	}

	return nil, fmt.Errorf("invalid response format")
}

// CreateProduct gọi API để tạo product mới
func (c *ProductHTTPClient) CreateProduct(ctx context.Context, product models.CreateProductInput) (*models.Product, error) {
	url := fmt.Sprintf("%s/products", c.baseURL)
	
	// Convert to API format
	apiProduct := models.Product{
		Name:              product.Name,
		Code:              product.Code,
		ShortCode:         product.ShortCode,
		ProductType:       product.ProductType,
		VAT:               product.VAT,
		VatPercent:        product.VatPercent,
		VatValue:          product.VatValue,
		CategoryID:        product.CategoryID,
		ProductGroupID:    product.ProductGroupID,
		PlatformID:        product.PlatformID,
		CanPreOrder:       product.CanPreOrder,
		ImageURL:          product.ImageURL,
		Note:              product.Note,
		MinOrderQuantity:  product.MinOrderQuantity,
		MedusaID:          product.MedusaID,
		OriginalCode:      product.OriginalCode,
		OriginalProductID: product.OriginalProductID,
	}

	jsonData, err := json.Marshal(apiProduct)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal product: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	var apiResponse models.APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert API response to our model
	if productData, ok := apiResponse.Data.(map[string]interface{}); ok {
		product := c.convertMapToProduct(productData)
		return &product, nil
	}

	return nil, fmt.Errorf("invalid response format")
}

// UpdateProduct gọi API để cập nhật product
func (c *ProductHTTPClient) UpdateProduct(ctx context.Context, id string, product models.UpdateProductInput) (*models.Product, error) {
	url := fmt.Sprintf("%s/products/%s", c.baseURL, id)
	
	// Convert to API format
	apiProduct := models.Product{}
	if product.Name != nil {
		apiProduct.Name = *product.Name
	}
	if product.Code != nil {
		apiProduct.Code = *product.Code
	}
	if product.ShortCode != nil {
		apiProduct.ShortCode = product.ShortCode
	}
	if product.ProductType != nil {
		apiProduct.ProductType = *product.ProductType
	}
	if product.VAT != nil {
		apiProduct.VAT = *product.VAT
	}
	if product.VatPercent != nil {
		apiProduct.VatPercent = *product.VatPercent
	}
	if product.VatValue != nil {
		apiProduct.VatValue = *product.VatValue
	}
	if product.CategoryID != nil {
		apiProduct.CategoryID = product.CategoryID
	}
	if product.ProductGroupID != nil {
		apiProduct.ProductGroupID = product.ProductGroupID
	}
	if product.PlatformID != nil {
		apiProduct.PlatformID = *product.PlatformID
	}
	if product.CanPreOrder != nil {
		apiProduct.CanPreOrder = *product.CanPreOrder
	}
	if product.ImageURL != nil {
		apiProduct.ImageURL = product.ImageURL
	}
	if product.Note != nil {
		apiProduct.Note = product.Note
	}
	if product.MinOrderQuantity != nil {
		apiProduct.MinOrderQuantity = *product.MinOrderQuantity
	}
	if product.MedusaID != nil {
		apiProduct.MedusaID = product.MedusaID
	}
	if product.OriginalCode != nil {
		apiProduct.OriginalCode = product.OriginalCode
	}
	if product.OriginalProductID != nil {
		apiProduct.OriginalProductID = product.OriginalProductID
	}

	jsonData, err := json.Marshal(apiProduct)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal product: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	// For update, we might need to fetch the updated product
	return c.GetProductByID(ctx, id)
}

// GetPaginatedProducts gọi API để lấy products có phân trang
func (c *ProductHTTPClient) GetPaginatedProducts(ctx context.Context, page, limit int, query string, filter models.ProductFilter) (*models.ProductPagination, error) {
	url := fmt.Sprintf("%s/products/paginate?page=%d&limit=%d", c.baseURL, page, limit)
	if query != "" {
		url += fmt.Sprintf("&q=%s", query)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var apiResponse models.APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert API response to our pagination model
	if data, ok := apiResponse.Data.(map[string]interface{}); ok {
		pagination := &models.ProductPagination{
			Page:  int32(data["page"].(float64)),
			Limit: int32(data["limit"].(float64)),
			Total: int32(data["total"].(float64)),
		}

		if items, ok := data["items"].([]interface{}); ok {
			for _, item := range items {
				if productData, ok := item.(map[string]interface{}); ok {
					product := c.convertMapToProduct(productData)
					pagination.Items = append(pagination.Items, &product)
				}
			}
		}

		return pagination, nil
	}

	return nil, fmt.Errorf("invalid response format")
}

// Helper function để convert map sang Product model
func (c *ProductHTTPClient) convertMapToProduct(data map[string]interface{}) models.Product {
	product := models.Product{}

	if id, ok := data["id"].(float64); ok {
		product.ID = uint64(id)
	}
	if name, ok := data["name"].(string); ok {
		product.Name = name
	}
	if code, ok := data["code"].(string); ok {
		product.Code = code
	}
	if shortCode, ok := data["short_code"].(string); ok {
		product.ShortCode = &shortCode
	}
	if productType, ok := data["product_type"].(string); ok {
		product.ProductType = productType
	}
	if vat, ok := data["vat"].(bool); ok {
		product.VAT = vat
	}
	if vatPercent, ok := data["vat_percent"].(float64); ok {
		product.VatPercent = int(vatPercent)
	}
	if vatValue, ok := data["vat_value"].(float64); ok {
		product.VatValue = int64(vatValue)
	}
	if categoryID, ok := data["category_id"].(float64); ok {
		catID := uint64(categoryID)
		product.CategoryID = &catID
	}
	if productGroupID, ok := data["product_group_id"].(float64); ok {
		groupID := uint64(productGroupID)
		product.ProductGroupID = &groupID
	}
	if platformID, ok := data["platform_id"].(float64); ok {
		product.PlatformID = uint64(platformID)
	}
	if canPreOrder, ok := data["can_pre_order"].(bool); ok {
		product.CanPreOrder = canPreOrder
	}
	if imageURL, ok := data["image_url"].(string); ok {
		product.ImageURL = &imageURL
	}
	if note, ok := data["note"].(string); ok {
		product.Note = &note
	}
	if minOrderQuantity, ok := data["min_order_quantity"].(float64); ok {
		product.MinOrderQuantity = int64(minOrderQuantity)
	}
	if medusaID, ok := data["medusa_id"].(string); ok {
		product.MedusaID = &medusaID
	}
	if originalCode, ok := data["original_code"].(string); ok {
		product.OriginalCode = &originalCode
	}
	if originalProductID, ok := data["original_product_id"].(float64); ok {
		origID := uint64(originalProductID)
		product.OriginalProductID = &origID
	}
	if searchTextV2, ok := data["search_text_v2"].(string); ok {
		product.SearchTextV2 = &searchTextV2
	}
	if searchText, ok := data["search_text"].(string); ok {
		product.SearchText = &searchText
	}
	if createdAt, ok := data["created_at"].(string); ok {
		product.CreatedAt = createdAt
	}
	if updatedAt, ok := data["updated_at"].(string); ok {
		product.UpdatedAt = updatedAt
	}

	return product
}
