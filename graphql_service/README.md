# GraphQL Service - Standalone

Đây là một GraphQL service riêng biệt chạy độc lập và gọi xuống HTTP API thay vì dùng chung database.

## Cấu trúc

```
graphql_service/
├── main.go                 # Entry point
├── go.mod                  # Go modules
├── gqlgen.yml             # GraphQL codegen config
├── graphql/
│   ├── schema.graphqls    # GraphQL schema
│   ├── resolver.go        # Resolver implementation
│   └── generated.go       # Generated code (sẽ được tạo)
├── http_client/
│   └── product_client.go  # HTTP client để gọi API
└── models/
    ├── models.go          # Custom models
    └── generated.go       # Generated models (sẽ được tạo)
```

## Cách hoạt động

1. **GraphQL Service** chạy trên port 8081 (hoặc GRAPHQL_PORT)
2. **HTTP API Service** chạy trên port 8080 (hoặc PORT)
3. GraphQL Service gọi xuống HTTP API thông qua HTTP client
4. Không có database connection trực tiếp từ GraphQL service

## Cài đặt và chạy

### 1. Tạo GraphQL code

```bash
cd graphql_service
go run github.com/99designs/gqlgen generate
```

### 2. Build và chạy

```bash
# Terminal 1: Chạy HTTP API service
cd /path/to/main/project
go run main.go

# Terminal 2: Chạy GraphQL service
cd graphql_service
go run main.go
```

### 3. Environment variables

Tạo file `.env` trong thư mục `graphql_service`:

```env
GRAPHQL_PORT=8081
API_BASE_URL=http://localhost:8080
```

## API Endpoints

### GraphQL Service (Port 8081)
- **GraphQL Playground**: http://localhost:8081
- **GraphQL Endpoint**: http://localhost:8081/query

### HTTP API Service (Port 8080)
- **GET /products**: Lấy tất cả products
- **GET /products/:id**: Lấy product theo ID
- **POST /products**: Tạo product mới
- **PUT /products/:id**: Cập nhật product
- **GET /products/paginate**: Lấy products có phân trang

## GraphQL Queries

### Lấy tất cả products
```graphql
query {
  products {
    id
    name
    description
    price
    category
    stock
    imageUrl
    createdAt
    updatedAt
  }
}
```

### Lấy product theo ID
```graphql
query {
  product(id: "1") {
    id
    name
    description
    price
    category
    stock
    imageUrl
  }
}
```

### Tạo product mới
```graphql
mutation {
  createProduct(input: {
    name: "Test Product"
    description: "Test Description"
    price: 99.99
    category: "electronics"
    stock: 10
    imageUrl: "https://example.com/image.jpg"
  }) {
    id
    name
    price
  }
}
```

### Cập nhật product
```graphql
mutation {
  updateProduct(id: "1", input: {
    name: "Updated Product"
    price: 149.99
    stock: 20
  }) {
    id
    name
    price
    stock
  }
}
```

### Lấy products có phân trang
```graphql
query {
  productsPaginated(page: 1, limit: 10) {
    page
    limit
    total
    items {
      id
      name
      price
      category
    }
  }
}
```

## Lợi ích của kiến trúc này

### 1. **Separation of Concerns**
- GraphQL service chỉ xử lý GraphQL logic
- HTTP API service xử lý business logic và database
- Không có dependency trực tiếp giữa GraphQL và database

### 2. **Scalability**
- Có thể scale GraphQL service và HTTP API service độc lập
- Có thể deploy GraphQL service ở nhiều region khác nhau

### 3. **Technology Flexibility**
- GraphQL service có thể được viết bằng ngôn ngữ khác
- HTTP API service có thể được thay thế bằng service khác

### 4. **Testing**
- Có thể test GraphQL service mà không cần database
- Có thể mock HTTP API responses

### 5. **Maintenance**
- Thay đổi GraphQL schema không ảnh hưởng HTTP API
- Thay đổi HTTP API không ảnh hưởng GraphQL service

## Development Workflow

### 1. Thêm GraphQL field mới
1. Cập nhật `graphql/schema.graphqls`
2. Chạy `go run github.com/99designs/gqlgen generate`
3. Implement resolver trong `graphql/resolver.go`
4. Thêm HTTP client method nếu cần

### 2. Thêm HTTP API endpoint mới
1. Thêm endpoint trong HTTP API service
2. Thêm method trong `http_client/product_client.go`
3. Cập nhật GraphQL resolver để sử dụng method mới

### 3. Testing
```bash
# Test HTTP API
curl http://localhost:8080/products

# Test GraphQL
curl -X POST http://localhost:8081/query \
  -H "Content-Type: application/json" \
  -d '{"query":"{ products { id name } }"}'
```

## Troubleshooting

### 1. GraphQL service không kết nối được HTTP API
- Kiểm tra `API_BASE_URL` trong `.env`
- Đảm bảo HTTP API service đang chạy
- Kiểm tra network connectivity

### 2. GraphQL playground không load
- Kiểm tra `GRAPHQL_PORT` trong `.env`
- Đảm bảo port không bị conflict
- Kiểm tra firewall settings

### 3. Generated code bị lỗi
- Chạy lại `go run github.com/99designs/gqlgen generate`
- Kiểm tra schema syntax
- Xóa file generated cũ và generate lại 