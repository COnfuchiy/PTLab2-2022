package domain

type Product struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

type IProductRepository interface {
	CountProducts() (int, error)
	FindProductById(id uint) (*Product, error)
	FindProductsByPagination(page, perPage int) ([]Product, error)
}

type IProductService interface {
	GetProductsCount() (int, error)
	GetProductById(id uint) (*Product, error)
	GetProducts(page, perPage int) ([]Product, error)
}
