# GraphQL + gRPC Architecture

## Flow Architecture
```
API Client -> GraphQL Server -> gRPC Server -> Database
```

## Cấu trúc Project
```
├── proto/                    # Protocol Buffer definitions
│   ├── product/             # Product service proto
│   │   ├── product.proto    # Product service proto
│   │   ├── product.pb.go    # Generated Go code
│   │   └── product_grpc.pb.go # Generated gRPC code
├── grpc/                    # gRPC implementation
│   └── server.go            # gRPC server
├── grpcclient/              # gRPC client
│   ├── client.go            # gRPC client
│   └── adapter.go           # Service adapter
├── graph/                   # GraphQL layer
│   ├── schema.graphqls      # GraphQL schema
│   ├── resolver.go          # Resolver interface
│   └── schema.resolvers.go  # Resolver implementation
├── graphql_server.go        # GraphQL server (direct service)
├── graphql_server_grpc.go   # GraphQL server (with gRPC client)
├── grpc_server              # Compiled gRPC server
├── graphql_server           # Compiled GraphQL server (direct)
├── graphql_server_grpc      # Compiled GraphQL server (with gRPC)
├── Makefile                 # Build and run commands
└── README_GRPC.md           # This file
```

## 🚀 Hướng dẫn chạy nhanh

### 1. Build tất cả services
```bash
make setup
```

### 2. Chạy cả hai server cùng lúc
```bash
make run-all
```

### 3. Hoặc chạy riêng lẻ

#### Terminal 1: Chạy gRPC Server
```bash
./grpc_server
```
**Output:**
```
gRPC server listening on :9090
```

#### Terminal 2: Chạy GraphQL Server (với gRPC client)
```bash
./graphql_server_grpc
```
**Output:**
```
GraphQL server listening on :8080
Connect to http://localhost:8080/ for GraphQL playground
```

## 📝 Test API

### 1. Test GraphQL Playground
Mở trình duyệt và truy cập: `http://localhost:8080`

### 2. Test với Postman

#### GraphQL Query - Lấy tất cả products
```
POST http://localhost:8080/query
Content-Type: application/json

{
  "query": "query { products { id name code productType sellingStatus } }"
}
```

#### GraphQL Query - Lấy product theo ID
```
POST http://localhost:8080/query
Content-Type: application/json

{
  "query": "query { product(id: \"1\") { id name code productType sellingStatus } }"
}
```

#### GraphQL Query - Lấy products với pagination
```
POST http://localhost:8080/query
Content-Type: application/json

{
  "query": "query { productsPaginated(page: 1, limit: 10) { products { id name code } total page limit } }"
}
```

#### GraphQL Mutation - Tạo product mới
```
POST http://localhost:8080/query
Content-Type: application/json

{
  "query": "mutation { createProduct(input: { name: \"Test Product\", code: \"TEST001\", productType: MATERIAL, type: \"GOODS\" }) { id name code } }"
}
```

#### GraphQL Mutation - Cập nhật product
```
POST http://localhost:8080/query
Content-Type: application/json

{
  "query": "mutation { updateProduct(id: \"1\", input: { name: \"Updated Product\" }) { id name code } }"
}
```

### 3. Test với cURL

#### Query products
```bash
curl -X POST http://localhost:8080/query \
  -H "Content-Type: application/json" \
  -d '{"query": "query { products { id name code } }"}'
```

#### Create product
```bash
curl -X POST http://localhost:8080/query \
  -H "Content-Type: application/json" \
  -d '{"query": "mutation { createProduct(input: { name: \"Test Product\", code: \"TEST001\", productType: MATERIAL, type: \"GOODS\" }) { id name code } }"}'
```

## 🔧 Các lệnh hữu ích

### Build riêng lẻ
```bash
# Build gRPC server
make build-grpc

# Build GraphQL server (direct service)
make build-graphql

# Build GraphQL server (with gRPC client)
make build-graphql-grpc
```

### Chạy riêng lẻ
```bash
# Chạy gRPC server
make run-grpc

# Chạy GraphQL server (direct service)
make run-graphql

# Chạy GraphQL server (with gRPC client)
make run-graphql-grpc
```

### Clean up
```bash
# Xóa các file build
make clean

# Regenerate proto files
make proto
```

## 📊 Monitoring

### Kiểm tra server status
```bash
# Kiểm tra gRPC server
netstat -tlnp | grep 9090

# Kiểm tra GraphQL server
netstat -tlnp | grep 8080
```

### Logs
- **gRPC Server logs:** Hiển thị business logic và database operations
- **GraphQL Server logs:** Hiển thị request/response và errors

## 🐛 Troubleshooting

### 1. Port đã được sử dụng
```bash
# Kill process sử dụng port 9090
sudo lsof -ti:9090 | xargs kill -9

# Kill process sử dụng port 8080
sudo lsof -ti:8080 | xargs kill -9
```

### 2. Database connection error
```bash
# Kiểm tra MySQL service
sudo systemctl status mysql

# Restart MySQL nếu cần
sudo systemctl restart mysql
```

### 3. Build error
```bash
# Clean và rebuild
make clean
make setup
```

## 🔍 Debug

### 1. Kiểm tra gRPC server
```bash
# Sử dụng grpcurl (nếu có)
grpcurl -plaintext localhost:9090 list
```

### 2. Kiểm tra GraphQL schema
```bash
# Truy cập GraphQL Playground
http://localhost:8080
```

### 3. Kiểm tra logs
```bash
# Theo dõi logs real-time
tail -f /var/log/syslog | grep -E "(grpc|graphql)"
```

## 📈 Performance

### 1. Load testing với Apache Bench
```bash
# Test GraphQL endpoint
ab -n 1000 -c 10 -p query.json -T application/json http://localhost:8080/query
```

### 2. Monitor resource usage
```bash
# Monitor CPU và Memory
htop

# Monitor network
iftop
```

## 🔐 Security

### 1. Production deployment
```bash
# Sử dụng HTTPS
# Sử dụng TLS cho gRPC
# Implement authentication
# Set up firewall rules
```

### 2. Environment variables
```bash
# Tạo file .env
cp .env.example .env

# Cấu hình database
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=your_database
```

## 📚 Next Steps

### 1. Thêm authentication
- JWT tokens
- API keys
- Role-based access control

### 2. Thêm caching
- Redis cache
- In-memory cache
- CDN for static files

### 3. Thêm monitoring
- Prometheus metrics
- Grafana dashboards
- Distributed tracing

### 4. Containerization
```dockerfile
# Dockerfile cho GraphQL server
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o graphql_server_grpc graphql_server_grpc.go
EXPOSE 8080
CMD ["./graphql_server_grpc"]
```

### 5. Kubernetes deployment
```yaml
# k8s-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: graphql-server
spec:
  replicas: 3
  selector:
    matchLabels:
      app: graphql-server
  template:
    metadata:
      labels:
        app: graphql-server
    spec:
      containers:
      - name: graphql-server
        image: your-registry/graphql-server:latest
        ports:
        - containerPort: 8080
```

## 🎯 Kết quả mong đợi

Sau khi chạy thành công, bạn sẽ có:

1. **gRPC Server** chạy trên port `9090`
2. **GraphQL Server** chạy trên port `8080`
3. **GraphQL Playground** tại `http://localhost:8080`
4. **API Client** có thể gọi GraphQL queries/mutations
5. **Flow hoàn chỉnh:** Client → GraphQL → gRPC → Database

**Chúc bạn thành công! 🚀** 