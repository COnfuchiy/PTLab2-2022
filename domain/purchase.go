package domain

type Purchase struct {
	ID        uint   `gorm:"primarykey"`
	Person    string `json:"person"`
	Address   string `json:"address"`
	Price     uint   `json:"price"`
	ProductID uint
	Product   Product
}

type IPurchaseRepository interface {
	InsertPurchase(purchase *Purchase) error
}

type IPurchaseService interface {
	CreatePurchase(purchase *Purchase) error
}
