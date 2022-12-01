package repository

import (
	"LAB2/domain"
)

type ProductDatabaseRepository struct {
	DBHandler DatabaseHandler
}

func NewProductDatabaseRepository(handler DatabaseHandler) ProductDatabaseRepository {
	return ProductDatabaseRepository{handler}
}

func (databaseRepository ProductDatabaseRepository) CountProducts() (int, error) {
	productCount, err := databaseRepository.DBHandler.CountProducts()
	if err != nil {
		return 0, err
	}
	return productCount, nil
}

func (databaseRepository ProductDatabaseRepository) FindProductById(id uint) (*domain.Product, error) {
	products, err := databaseRepository.DBHandler.FindProductById(id)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (databaseRepository ProductDatabaseRepository) FindProductsByPagination(page, perPage int) ([]domain.Product, error) {
	products, err := databaseRepository.DBHandler.FindProductsByPagination(page, perPage)
	if err != nil {
		return nil, err
	}
	return products, nil
}
