package db

import (
	"LAB2/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseHandler struct {
	DB *gorm.DB
}

func NewSqliteDatabaseHandler(databaseFileName string) (*DatabaseHandler, error) {
	dbHandler, err := gorm.Open(sqlite.Open(databaseFileName), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if res := dbHandler.Exec("PRAGMA foreign_keys = ON", nil); res.Error != nil {
		return nil, err
	}

	err = dbHandler.AutoMigrate(&domain.Product{}, &domain.Purchase{})
	if err != nil {
		return nil, err
	}
	return &DatabaseHandler{dbHandler}, nil
}

func (handler DatabaseHandler) InsertPurchase(purchase *domain.Purchase) error {
	result := handler.DB.Model(&domain.Purchase{}).Create(purchase)
	return result.Error
}
func (handler DatabaseHandler) CountProducts() (int, error) {
	var totalCount int64 = 0
	result := handler.DB.Model(&domain.Product{}).Count(&totalCount)
	return int(totalCount), result.Error
}
func (handler DatabaseHandler) FindProductById(id uint) (*domain.Product, error) {
	var product domain.Product
	result := handler.DB.Model(&domain.Product{}).Where("id = ?", id).First(&product)
	return &product, result.Error

}
func (handler DatabaseHandler) FindProductsByPagination(page, perPage int) ([]domain.Product, error) {
	var products []domain.Product
	offset := (page - 1) * perPage
	result := handler.DB.Offset(offset).Limit(perPage).Find(&products)
	return products, result.Error
}
