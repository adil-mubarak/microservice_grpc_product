package service

import (
	"context"
	"fmt"
	"microservice_grpc_product/models"
	product "microservice_grpc_product/pb/product"

	"gorm.io/gorm"
)

type ProductServiceServer struct {
	product.UnimplementedProductServiceServer
	DB *gorm.DB
}

func (s *ProductServiceServer) CreateProduct(ctx context.Context, req *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	products := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       float64(req.Price),
		Stock:       req.Stock,
	}

	if err := s.DB.Create(&products).Error; err != nil {
		return &product.CreateProductResponse{
			Message: "Falied to create product",
		}, err
	}

	return &product.CreateProductResponse{
		Product: &product.Product{
			Id:          uint32(products.ID),
			Name:        products.Name,
			Description: products.Description,
			Price:       float32(products.Price),
			Stock:       products.Stock,
		},
		Message: "Product created successfully",
	}, nil
}

func (s *ProductServiceServer) GetProduct(ctx context.Context, req *product.GetProductRequest) (*product.GetProductResponse, error) {
	var products models.Product

	if err := s.DB.First(&products, req.GetId()).Error; err != nil {
		return &product.GetProductResponse{
			Message: fmt.Sprintf("Product with ID %d not found", req.GetId()),
		}, err
	}

	return &product.GetProductResponse{
		Product: &product.Product{
			Id:          uint32(products.ID),
			Name:        products.Name,
			Description: products.Description,
			Price:       float32(products.Price),
			Stock:       products.Stock,
		},
		Message: "Product retrived successfully",
	}, nil
}

func (s *ProductServiceServer) UpdateProduct(ctx context.Context, req *product.UpdateProductRequest) (*product.UpdateProductResponse, error) {
	var products models.Product

	if err := s.DB.First(&products,req.GetId()).Error; err != nil{
		return &product.UpdateProductResponse{
			Message: fmt.Sprintf("Product with ID %d is not found",req.GetId()),
		},err
	}

	products.Name = req.GetName()
	products.Description = req.GetDescription()
	products.Price = float64(req.GetPrice())
	products.Stock = req.GetStock()

	if err := s.DB.Save(&products).Error; err != nil{
		return &product.UpdateProductResponse{
			Message: "Failed to update product",
		},err
	}

	return &product.UpdateProductResponse{
		Product: &product.Product{
			Id: uint32(products.ID),
			Name: products.Name,
			Description: products.Description,
			Price: float32(products.Price),
			Stock: products.Stock,
		},
		Message: "Prduct updated successfully",
	},nil
}


func (s *ProductServiceServer) DeleteProduct(ctx context.Context, req *product.DeleteProductRequest)(*product.DeleteProductResponse,error){
	var products models.Product

	if err := s.DB.Delete(&products,req.GetId()).Error; err != nil{
		return &product.DeleteProductResponse{
			Message: fmt.Sprintf("Failed ro delete product with ID %d ",req.GetId()),
		},nil
	}

	return &product.DeleteProductResponse{
		Message: "Product deleted successfully",
	},nil
}

func (s *ProductServiceServer)GetAllProduts(ctx context.Context,req *product.GetAllProductsRequest)(*product.GetAllProductsResponse,error){
	var products []models.Product

	if err := s.DB.Find(&products).Error; err != nil{
		return &product.GetAllProductsResponse{
			Product: nil,
		},err
	}

	var pbProducts []*product.Product
	for _, p := range products{
		pbProducts = append(pbProducts, &product.Product{
			Id: uint32(p.ID),
			Name: p.Name,
			Description: p.Description,
			Price: float32(p.Price),
			Stock: p.Stock,
		})
	}

	return &product.GetAllProductsResponse{
		Product: pbProducts,
	},nil

}	