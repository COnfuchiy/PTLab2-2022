package domain

import (
	"time"
)

type Purchase struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Person    string `json:"person"`
	Address   string `json:"address"`
	Price     uint   `json:"price"`
	CreatedAt time.Time
	ProductID uint `json:"product_id"`
	Product   Product
}

type PurchaseRepository interface {
	InsertPurchase(purchase *Purchase) error
}
