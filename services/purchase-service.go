package services

import (
	"LAB2/domain"
	"log"
)

type PurchaseService struct {
	PurchaseRepository         domain.IPurchaseRepository
	productPurchasesToDiscount int
}

func NewPurchaseService(repository domain.IPurchaseRepository) domain.IPurchaseService {
	return &PurchaseService{repository, 10}
}

func (service *PurchaseService) CreatePurchase(purchase *domain.Purchase) error {
	err := service.PurchaseRepository.InsertPurchase(purchase)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (service *PurchaseService) CheckProductWillHaveDiscount(product *domain.Product) (bool, error) {
	productPurchasesCount, err := service.PurchaseRepository.CountPurchasesByProductId(product.ID)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	if product.Discount.Percent != 0 {
		return false, nil
	}
	if productPurchasesCount == service.productPurchasesToDiscount {
		return true, nil
	}
	return false, nil
}
