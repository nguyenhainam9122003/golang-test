package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"test/core/domain"
	"test/core/model"
)

type MockProductRepository struct {
	mock.Mock
	domain.ProductRepository
}

func (m *MockProductRepository) Create(ctx context.Context, product *model.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}
func (m *MockProductRepository) FindByID(ctx context.Context, id uint) (*model.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Product), args.Error(1)
}
func (m *MockProductRepository) FetchPaginated(ctx context.Context, limit, offset int, query string, filter model.ProductFilter) ([]model.Product, error) {
	args := m.Called(ctx, limit, offset, query, filter)
	return args.Get(0).([]model.Product), args.Error(1)
}

func TestProductService_Create(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()
	product := &model.Product{ID: 1, Name: "Test Product"}
	mockRepo.On("Create", ctx, product).Return(nil)
	err := service.Create(ctx, product)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetByID(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()
	product := &model.Product{ID: 1, Name: "A"}
	mockRepo.On("FindByID", ctx, uint(1)).Return(product, nil)
	result, err := service.GetByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, product, result)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetPaginated(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()
	products := []model.Product{{ID: 1, Name: "A"}}
	filter := model.ProductFilter{}
	mockRepo.On("FetchPaginated", ctx, 10, 0, "", filter).Return(products, nil)
	result, err := service.GetPaginated(ctx, 10, 0, "", filter)
	assert.NoError(t, err)
	assert.Equal(t, products, result)
	mockRepo.AssertExpectations(t)
} 