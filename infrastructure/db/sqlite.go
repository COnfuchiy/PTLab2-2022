package db

import (
	"LAB2/domain"
	"LAB2/interface/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseHandler struct {
	DB *gorm.DB
}

func NewSqliteDatabaseHandler(databaseFileName string) (repository.IDatabaseHandler, error) {
	dbHandler, err := gorm.Open(sqlite.Open(databaseFileName), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if res := dbHandler.Exec("PRAGMA foreign_keys = ON", nil); res.Error != nil {
		return nil, err
	}

	err = dbHandler.AutoMigrate(&domain.Product{}, &domain.Purchase{}, &domain.Discount{})
	if err != nil {
		return nil, err
	}
	return &DatabaseHandler{dbHandler}, nil
}

func (handler DatabaseHandler) InsertPurchase(purchase *domain.Purchase) error {
	return handler.DB.Model(&domain.Purchase{}).Create(&purchase).Error
}
func (handler DatabaseHandler) CountProducts() (int, error) {
	var totalCount int64 = 0
	result := handler.DB.Model(&domain.Product{}).Count(&totalCount)
	return int(totalCount), result.Error
}
func (handler DatabaseHandler) FindProductById(id uint) (*domain.Product, error) {
	var product domain.Product
	result := handler.DB.Model(&domain.Product{}).Preload("Discount").Where("id = ?", id).First(&product)
	return &product, result.Error

}
func (handler DatabaseHandler) FindProductsByPagination(page, perPage int) ([]domain.Product, error) {
	var products []domain.Product
	offset := (page - 1) * perPage
	result := handler.DB.Model(&domain.Product{}).Preload("Discount").Offset(offset).Limit(perPage).Find(&products)
	return products, result.Error
}

func (handler DatabaseHandler) UpdateProduct(product *domain.Product) error {
	return handler.DB.Model(&domain.Product{}).Where("id = ?", product.ID).Updates(&product).Error
}

func (handler DatabaseHandler) CountPurchasesByProductId(productId uint) (int, error) {
	var purchasesCount int64
	result := handler.DB.Model(&domain.Purchase{}).Where("product_id = ?", productId).Count(&purchasesCount)
	return int(purchasesCount), result.Error
}

func (handler DatabaseHandler) InsertDiscount(discount *domain.Discount) error {
	return handler.DB.Model(&domain.Discount{}).Create(&discount).Error
}
func (handler DatabaseHandler) DeleteDiscount(discount *domain.Discount) error {
	return handler.DB.Model(&domain.Discount{}).Where("id = ?", discount.ID).Delete(&discount).Error
}
