package services

import (
	"LAB2/domain"
	"log"
)

type ProductService struct {
	ProductRepository domain.IProductRepository
}

func NewProductService(repository domain.IProductRepository) domain.IProductService {
	return &ProductService{repository}
}

func (service *ProductService) GetProductsCount() (int, error) {
	productCount, err := service.ProductRepository.CountProducts()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return productCount, nil
}

func (service *ProductService) GetProductById(id uint) (*domain.Product, error) {
	product, err := service.ProductRepository.FindProductById(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return product, nil
}

func (service *ProductService) GetProducts(page, perPage int) ([]domain.Product, error) {
	products, err := service.ProductRepository.FindProductsByPagination(page, perPage)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return products, nil
}
