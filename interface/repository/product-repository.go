package repository

import (
	"LAB2/domain"
)

type ProductRepository struct {
	DBHandler DatabaseHandler
}

func NewProductRepository(handler DatabaseHandler) ProductRepository {
	return ProductRepository{handler}
}

func (repository ProductRepository) CountProducts() (int, error) {
	productCount, err := repository.DBHandler.CountProducts()
	if err != nil {
		return 0, err
	}
	return productCount, nil
}

func (repository ProductRepository) FindProductById(id uint) (*domain.Product, error) {
	products, err := repository.DBHandler.FindProductById(id)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (repository ProductRepository) FindProductsByPagination(page, perPage int) ([]domain.Product, error) {
	products, err := repository.DBHandler.FindProductsByPagination(page, perPage)
	if err != nil {
		return nil, err
	}
	return products, nil
}
