package repository

import (
	"LAB2/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&domain.Product{}, &domain.Purchase{})
	if err != nil {
		return
	}

	DB = database
}
