package mocks

import (
	"LAB2/domain"
	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (mock *MockProductService) GetProductsCount() (int, error) {
	args := mock.Called()
	totalCount := args.Int(0)
	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}
	return totalCount, err
}

func (mock *MockProductService) GetProductById(id uint) (*domain.Product, error) {
	args := mock.Called(id)
	var product *domain.Product
	if args.Get(0) != nil {
		product = args.Get(0).(*domain.Product)
	}
	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}
	return product, err
}

func (mock *MockProductService) GetProducts(page, perPage int) ([]domain.Product, error) {
	args := mock.Called(page, perPage)
	var products []domain.Product
	if args.Get(0) != nil {
		products = args.Get(0).([]domain.Product)
	}
	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}
	return products, err
}
