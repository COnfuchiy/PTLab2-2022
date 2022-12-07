package repository

import (
	"LAB2/domain"
)

type ProductRepository struct {
	DBHandler IDatabaseHandler
}

func NewProductRepository(handler IDatabaseHandler) domain.IProductRepository {
	return ProductRepository{handler}
}

func (repository ProductRepository) CountProducts() (int, error) {
	return repository.DBHandler.CountProducts()
}

func (repository ProductRepository) FindProductById(id uint) (*domain.Product, error) {
	return repository.DBHandler.FindProductById(id)
}

func (repository ProductRepository) FindProductsByPagination(page, perPage int) ([]domain.Product, error) {
	return repository.DBHandler.FindProductsByPagination(page, perPage)
}

func (repository ProductRepository) UpdateProduct(product *domain.Product) error {
	return repository.DBHandler.UpdateProduct(product)
}
