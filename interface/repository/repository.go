package repository

import (
	"LAB2/domain"
)

type IDatabaseHandler interface {
	UpdateProduct(product *domain.Product) error
	InsertPurchase(purchase *domain.Purchase) error
	CountProducts() (int, error)
	FindProductById(id uint) (*domain.Product, error)
	FindProductsByPagination(page, perPage int) ([]domain.Product, error)
	CountPurchasesByProductId(productId uint) (int, error)
	InsertDiscount(discount *domain.Discount) error
	DeleteDiscount(discount *domain.Discount) error
}
