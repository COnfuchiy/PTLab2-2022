package repository

import (
	"LAB2/domain"
)

type DatabaseHandler interface {
	InsertPurchase(purchase *domain.Purchase) error
	CountProducts() (int, error)
	FindProductById(id uint) (*domain.Product, error)
	FindProductsByPagination(page, perPage int) ([]domain.Product, error)
}
