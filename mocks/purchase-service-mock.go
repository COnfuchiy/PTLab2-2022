package mocks

import (
	"LAB2/domain"
	"github.com/stretchr/testify/mock"
)

type MockPurchaseService struct {
	mock.Mock
}

func (mock *MockPurchaseService) CreatePurchase(purchase *domain.Purchase) error {
	args := mock.Called(purchase)
	var err error
	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}
	return err
}

func (mock *MockPurchaseService) CheckProductWillHaveDiscount(product *domain.Product) (bool, error) {
	args := mock.Called(product)
	productWillHaveDiscount := args.Bool(0)
	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}
	return productWillHaveDiscount, err
}
