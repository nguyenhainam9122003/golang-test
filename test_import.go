package main

import (
	"fmt"
	grpcproduct "test/proto/product"
)

func main() {
	fmt.Println("Testing import...")
	fmt.Printf("Product type: %T\n", grpcproduct.Product{})
} 