package mocks

import (
	"LAB2/domain"
	"github.com/stretchr/testify/mock"
)

type MockDiscountService struct {
	mock.Mock
}

func (mock *MockDiscountService) CreateDiscount(discount *domain.Discount) error {
	args := mock.Called(discount)
	var err error
	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}
	return err
}
func (mock *MockDiscountService) DeleteDiscount(discount *domain.Discount) error {
	args := mock.Called(discount)
	var err error
	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}
	return err
}
func (mock *MockDiscountService) CheckOrDeleteProductDiscount(product domain.Product) (domain.Product, error) {
	args := mock.Called(product)
	var returnedProduct domain.Product
	if args.Get(0) != nil {
		returnedProduct = args.Get(0).(domain.Product)
	}
	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}
	return returnedProduct, err
}
func (mock *MockDiscountService) CheckDateOrDeleteDiscount(discount domain.Discount) (domain.Discount, error) {
	args := mock.Called(discount)
	var returnedDiscount domain.Discount
	if args.Get(0) != nil {
		returnedDiscount = args.Get(0).(domain.Discount)
	}
	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}
	return returnedDiscount, err
}
func (mock *MockDiscountService) GetPriceWithDiscount(price, percent uint) uint {
	args := mock.Called(price, percent)
	var priceWithDiscount uint
	if args.Get(0) != nil {
		priceWithDiscount = args.Get(0).(uint)
	}
	return priceWithDiscount
}
