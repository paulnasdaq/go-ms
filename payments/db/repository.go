package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	Transactions TransactionRepository
}

func NewRepository() (*Repository, error) {
	db, err := gorm.Open(sqlite.Open("transactions.db"), &gorm.Config{})
	if err != nil {
		log.Println("Failed to initialize repository: ", err)
		return nil, err
	}
	if err := db.AutoMigrate(&Transaction{}); err != nil {
		log.Println("Failed to initialize repository: ", err)
		return nil, err
	}
	return &Repository{Transactions: TransactionRepository{db: db}}, nil
}
