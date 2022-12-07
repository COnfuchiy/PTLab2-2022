package repository

import (
	"LAB2/domain"
)

type DiscountRepository struct {
	DBHandler IDatabaseHandler
}

func NewDiscountRepository(handler IDatabaseHandler) domain.IDiscountRepository {
	return DiscountRepository{handler}
}

func (repository DiscountRepository) InsertDiscount(discount *domain.Discount) error {
	return repository.DBHandler.InsertDiscount(discount)
}

func (repository DiscountRepository) DeleteDiscount(discount *domain.Discount) error {
	return repository.DBHandler.DeleteDiscount(discount)
}
