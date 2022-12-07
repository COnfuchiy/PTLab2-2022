package repository

import (
	"LAB2/domain"
)

type PurchaseRepository struct {
	DBHandler IDatabaseHandler
}

func NewPurchaseRepository(handler IDatabaseHandler) domain.IPurchaseRepository {
	return PurchaseRepository{handler}
}

func (repository PurchaseRepository) InsertPurchase(purchase *domain.Purchase) error {
	return repository.DBHandler.InsertPurchase(purchase)
}

func (repository PurchaseRepository) CountPurchasesByProductId(productId uint) (int, error) {
	return repository.DBHandler.CountPurchasesByProductId(productId)
}
