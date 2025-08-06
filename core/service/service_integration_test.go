package service

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"test/core/model"
	"test/core/repository"
)

func cleanupTestDB(db *gorm.DB) {
	db.Exec("DELETE FROM product_prices WHERE product_id IN (SELECT id FROM products WHERE code LIKE 'TEST-%')")
	db.Exec("DELETE FROM products WHERE code LIKE 'TEST-%'")
}

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:passwordhere@tcp(localhost:13306)/fulfillment?parseTime=true"
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&model.Product{}, &model.ProductPrice{})
	return db
}

func TestProductService_Integration_CreateAndGetByID(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(db)
	repo := repository.NewProductRepository(db)
	service := NewProductService(repo)
	ctx := context.Background()

	product := &model.Product{
		Name:  "Integration Product",
		Code:  "TEST-INT-001",
		Type:  "TYPE1",
		IsActive: true,
		ProductType: "MATERIAL",
		SellingStatus: "AVAILABLE",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := service.Create(ctx, product)
	assert.NoError(t, err)
	assert.NotZero(t, product.ID)

	fetched, err := service.GetByID(ctx, uint(product.ID))
	assert.NoError(t, err)
	assert.Equal(t, product.Name, fetched.Name)
	assert.Equal(t, product.Code, fetched.Code)
}

func TestProductService_Integration_GetPaginated(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(db)
	repo := repository.NewProductRepository(db)
	service := NewProductService(repo)
	ctx := context.Background()

	// Insert multiple products
	for i := 0; i < 3; i++ {
		p := &model.Product{
			Name:  "Paginate Product " + string(rune('A'+i)),
			Code:  "TEST-PAGE-00" + string(rune('1'+i)),
			Type:  "TYPE1",
			IsActive: true,
			ProductType: "MATERIAL",
			SellingStatus: "AVAILABLE",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		_ = service.Create(ctx, p)
	}

	filter := model.ProductFilter{}
	products, err := service.GetPaginated(ctx, 2, 0, "", filter)
	assert.NoError(t, err)
	assert.True(t, len(products) >= 2)
} 