package actions

import (
	"LAB2/domain"
	"log"
)

type PurchaseInteractor struct {
	PurchaseRepository domain.PurchaseRepository
}

func NewPurchaseInteractor(repository domain.PurchaseRepository) PurchaseInteractor {
	return PurchaseInteractor{repository}
}

func (interactor PurchaseInteractor) Insert(purchase *domain.Purchase) error {
	err := interactor.PurchaseRepository.InsertPurchase(purchase)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
