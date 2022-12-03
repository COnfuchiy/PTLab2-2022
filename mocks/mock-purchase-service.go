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
