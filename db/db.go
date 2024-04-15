package db

import (
	"github.com/rahulii/Library/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Init(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Book{})
	if err != nil {
		return nil, err
	}

	return db, nil
}