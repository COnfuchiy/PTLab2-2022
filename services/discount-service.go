package services

import (
	"LAB2/domain"
	"log"
	"time"
)

type DiscountService struct {
	DiscountRepository domain.IDiscountRepository
}

func NewDiscountService(repository domain.IDiscountRepository) domain.IDiscountService {
	return &DiscountService{repository}
}

func (service *DiscountService) CreateDiscount(discount *domain.Discount) error {
	err := service.DiscountRepository.InsertDiscount(discount)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (service *DiscountService) DeleteDiscount(discount *domain.Discount) error {
	err := service.DiscountRepository.DeleteDiscount(discount)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (service *DiscountService) CheckOrDeleteProductDiscount(product domain.Product) (domain.Product, error) {
	if time.Now().After(product.Discount.EndDate) {
		err := service.DeleteDiscount(&product.Discount)
		if err != nil {
			return product, err
		}
		product.Discount = domain.Discount{}
	}
	return product, nil
}

func (service *DiscountService) CheckDateOrDeleteDiscount(discount domain.Discount) (domain.Discount, error) {
	if time.Now().After(discount.EndDate) {
		err := service.DeleteDiscount(&discount)
		if err != nil {
			return discount, err
		}
		return domain.Discount{}, nil
	}
	return discount, nil
}

func (service *DiscountService) GetPriceWithDiscount(price, percent uint) uint {
	if percent != 0 {
		outputPrice := float64(price) * (float64(100-percent) / 100)
		if outputPrice > 0 {
			return uint(outputPrice)
		}
	}
	return price
}
