package mocks

import (
	"LAB2/domain"
	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
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

func (mock *MockProductService) GetProducts(page int) ([]domain.Product, error) {
	args := mock.Called(page)
	var products []domain.Product
	if args.Get(0) != nil {
		products = args.Get(0).([]domain.Product)
		//for _, product := range products {
		//	println(product.ID)
		//}
	}
	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}
	return products, err
}
func (mock *MockProductService) GetPaginationInfo(page int) (int, int, error) {
	args := mock.Called(page)
	mockedPage := args.Int(0)
	totalPagesCount := args.Int(1)
	var err error
	if args.Get(2) != nil {
		err = args.Get(2).(error)
	}
	return mockedPage, totalPagesCount, err
}
func (mock *MockProductService) DecreaseProductCount(product *domain.Product) error {
	args := mock.Called(product)
	product.Price -= 1
	var err error
	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}
	return err
}
