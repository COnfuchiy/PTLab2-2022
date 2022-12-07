package services

import (
	"LAB2/domain"
	"log"
	"math"
)

type ProductService struct {
	productRepository domain.IProductRepository
	pageSize          int
}

func NewProductService(repository domain.IProductRepository) domain.IProductService {
	return &ProductService{repository, 10}
}

func (service *ProductService) GetProductById(id uint) (*domain.Product, error) {
	product, err := service.productRepository.FindProductById(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return product, nil
}

func (service *ProductService) GetProducts(page int) ([]domain.Product, error) {
	products, err := service.productRepository.FindProductsByPagination(page, service.pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return products, nil
}

func (service *ProductService) GetPaginationInfo(page int) (int, int, error) {
	totalCount, err := service.productRepository.CountProducts()
	if err != nil {
		log.Println(err.Error())
		return 0, 0, err
	}
	totalPagesCount := int(math.Ceil(float64(totalCount) / float64(service.pageSize)))
	if page > totalPagesCount {
		return totalPagesCount, totalCount, nil
	}
	return page, totalPagesCount, nil
}

func (service *ProductService) DecreaseProductCount(product *domain.Product) error {
	product.Count -= 1
	err := service.productRepository.UpdateProduct(product)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
