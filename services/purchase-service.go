package services

import (
	"LAB2/domain"
	"log"
)

type PurchaseService struct {
	PurchaseRepository domain.IPurchaseRepository
}

func NewPurchaseService(repository domain.IPurchaseRepository) domain.IPurchaseService {
	return &PurchaseService{repository}
}

func (service *PurchaseService) CreatePurchase(purchase *domain.Purchase) error {
	err := service.PurchaseRepository.InsertPurchase(purchase)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
