package graphql

import (
	"test/graphql_service/http_client"
)

// Resolver cho GraphQL service
type Resolver struct {
	ProductHTTPClient *http_client.ProductHTTPClient
}

// Resolver implementation sẽ được tách ra thành các file riêng:
// - mutation.resolvers.go: Chứa các mutation resolvers
// - query.resolvers.go: Chứa các query resolvers

// Resolver types - sẽ được generate bởi gqlgen
// type mutationResolver struct{ *Resolver }
// type queryResolver struct{ *Resolver }

// func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }
// func (r *Resolver) Query() QueryResolver       { return &queryResolver{r} }

// Interface definitions - sẽ được generate bởi gqlgen
// type MutationResolver interface {
// 	CreateProduct(ctx context.Context, input models.CreateProductInput) (*models.Product, error)
// 	UpdateProduct(ctx context.Context, id string, input models.UpdateProductInput) (*models.Product, error)
// 	DeleteProduct(ctx context.Context, id string) (bool, error)
// }

// type QueryResolver interface {
// 	Products(ctx context.Context) ([]*models.Product, error)
// 	Product(ctx context.Context, id string) (*models.Product, error)
// 	GetMaterialProduct(ctx context.Context) ([]*models.MaterialProduct, error)
// 	ProductsPaginated(ctx context.Context, page *int32, limit *int32, query *string, filter *models.ProductFilter) (*models.ProductPagination, error)
// }
