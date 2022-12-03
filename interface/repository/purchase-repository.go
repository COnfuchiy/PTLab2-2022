package repository

import (
	"LAB2/domain"
)

type PurchaseRepository struct {
	DBHandler DatabaseHandler
}

func NewPurchaseRepository(handler DatabaseHandler) PurchaseRepository {
	return PurchaseRepository{handler}
}

func (repository PurchaseRepository) InsertPurchase(purchase *domain.Purchase) error {
	err := repository.DBHandler.InsertPurchase(purchase)
	if err != nil {
		return err
	}
	return nil
}
