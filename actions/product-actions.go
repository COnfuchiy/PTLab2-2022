package actions

import (
	"LAB2/domain"
	"log"
)

type ProductInteractor struct {
	ProductRepository domain.ProductRepository
}

func NewProductInteractor(repository domain.ProductRepository) ProductInteractor {
	return ProductInteractor{repository}
}

func (interactor *ProductInteractor) CountProducts() (int, error) {
	productCount, err := interactor.ProductRepository.CountProducts()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return productCount, nil
}

func (interactor *ProductInteractor) FindProductById(id uint) (*domain.Product, error) {
	product, err := interactor.ProductRepository.FindProductById(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return product, nil
}

func (interactor *ProductInteractor) FindProductsByPagination(page, perPage int) ([]domain.Product, error) {
	products, err := interactor.ProductRepository.FindProductsByPagination(page, perPage)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return products, nil
}
