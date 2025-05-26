package db

import (
	"github.com/google/uuid"
	"github.com/paulnasdaq/fms-v2/common"
	"gorm.io/gorm"
	"log"
)

type TransactionType string

const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"
)

type Transaction struct {
	gorm.Model
	ID     uuid.UUID
	Type   TransactionType `gorm:"type:varchar"`
	Amount float64
	UserID uuid.UUID
}

type TransactionRepository struct {
	db      *gorm.DB
	clients common.ClientManager
}

func (r *TransactionRepository) Add(transactionType TransactionType, amount float64, userID uuid.UUID) (*Transaction, error) {
	newTransaction := Transaction{
		ID:     uuid.New(),
		Type:   transactionType,
		Amount: amount,
		UserID: userID,
	}
	if res := r.db.Create(&newTransaction); res.Error != nil {
		log.Println(res.Error)
		return nil, res.Error
	}
	return &newTransaction, nil
}

func (r *TransactionRepository) Get(ID uuid.UUID) (*Transaction, error) {
	transaction := Transaction{}
	if res := r.db.First(&transaction, ID); res.Error != nil {
		log.Println(res.Error)
		return nil, res.Error
	}
	return &transaction, nil
}

type balance struct {
	Balance float64
}

func (r *TransactionRepository) Balance(userID uuid.UUID) (float64, error) {
	balance := balance{}
	res := r.db.Raw(`SELECT SUM(CASE WHEN type = 'credit' THEN amount ELSE 0 END) - SUM(CASE WHEN type = 'debit' THEN amount ELSE 0 END) AS balance from transactions where user_id = ?`, userID).Scan(&balance)
	if res.Error != nil {
		log.Println("Failed to get balance", res.Error)
		return 0, res.Error
	}
	return balance.Balance, nil
}
