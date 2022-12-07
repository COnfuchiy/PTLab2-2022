package domain

type Product struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `json:"name"`
	Count    int    `json:"count"`
	Price    uint   `json:"price"`
	Discount Discount
}

type IProductRepository interface {
	UpdateProduct(product *Product) error
	CountProducts() (int, error)
	FindProductById(id uint) (*Product, error)
	FindProductsByPagination(page, perPage int) ([]Product, error)
}

type IProductService interface {
	GetProductById(id uint) (*Product, error)
	GetProducts(page int) ([]Product, error)
	GetPaginationInfo(page int) (int, int, error)
	DecreaseProductCount(product *Product) error
}
