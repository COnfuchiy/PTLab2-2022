package repository

import (
	"LAB2/domain"
)

type PurchaseDatabaseRepository struct {
	DBHandler DatabaseHandler
}

func NewPurchaseDatabaseRepository(handler DatabaseHandler) PurchaseDatabaseRepository {
	return PurchaseDatabaseRepository{handler}
}

func (databaseRepository PurchaseDatabaseRepository) InsertPurchase(purchase *domain.Purchase) error {
	err := databaseRepository.DBHandler.InsertPurchase(purchase)
	if err != nil {
		return err
	}
	return nil
}
