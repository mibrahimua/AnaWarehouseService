package service

import (
	"AnaProductService/model"
	"AnaProductService/repository"
)

type ProductService struct {
	productRepository *repository.ProductRepository
}

func NewProductService(productRepository *repository.ProductRepository) *ProductService {
	return &ProductService{productRepository}
}

func (us *ProductService) GetListProductByName(productName string) ([]model.Product, error) {
	return us.productRepository.GetListProductByName(productName)
}
