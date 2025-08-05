module test/graphql_service

go 1.23.0

replace test/proto => ../proto

toolchain go1.24.5

require (
	github.com/99designs/gqlgen v0.17.78
	github.com/joho/godotenv v1.5.1
	github.com/vektah/gqlparser/v2 v2.5.30
	google.golang.org/grpc v1.74.2
	test/proto v0.0.0-00010101000000-000000000000
)

require (
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250528174236-200df99c418a // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)
