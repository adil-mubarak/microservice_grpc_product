package main

import (
	"fmt"
	"log"
	"microservice_grpc_product/db"
	"microservice_grpc_product/pb/product"
	"microservice_grpc_product/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := db.ConnectDatebase()
	if err != nil{
		log.Fatalf("Failed to connect to database: %v",err)
	}

	grpcServer := grpc.NewServer()

	productService := &service.ProductServiceServer{
		DB: db,
	}
	product.RegisterProductServiceServer(grpcServer,productService)

	lis, err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatalf("Failed to listen on port 8080: %v",err)
	}

	fmt.Println("Server is running on port :8080")
	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Failed to serve gRPC server: %v",err)
	}
}