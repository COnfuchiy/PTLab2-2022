package domain

import (
	"time"
)

var DiscountDefaultPercent uint = 15

var DiscountDefaultDaysInterval = 14

type Discount struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Percent   uint      `json:"percent"`
	EndDate   time.Time `json:"end_date"`
	ProductID uint      `json:"product_id"`
}

type IDiscountRepository interface {
	InsertDiscount(discount *Discount) error
	DeleteDiscount(discount *Discount) error
}

type IDiscountService interface {
	CreateDiscount(discount *Discount) error
	DeleteDiscount(discount *Discount) error
	CheckDateOrDeleteDiscount(discount Discount) (Discount, error)
	GetPriceWithDiscount(price, percent uint) uint
	CheckOrDeleteProductDiscount(product Product) (Product, error)
}

func NewDiscount(productId, percent uint, durationDays int) Discount {
	return Discount{
		Percent:   percent,
		EndDate:   time.Now().AddDate(0, 0, durationDays),
		ProductID: productId,
	}
}

func NewDefaultDiscount(productId uint) Discount {
	return NewDiscount(productId, DiscountDefaultPercent, DiscountDefaultDaysInterval)
}
