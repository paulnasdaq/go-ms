package services

import (
	"github.com/google/uuid"
	"github.com/paulnasdaq/fms-v2/payments/db"
	users "github.com/paulnasdaq/fms-v2/users"
	"log"
	"time"
)

type Transaction struct {
	ID        string
	Amount    float64
	CreatedAt time.Time
	UserID    string
	User      users.User
	Type      db.TransactionType
}

type TransactionService struct {
	Repository *db.Repository
}

func (s *TransactionService) Add(transactionType db.TransactionType, amount float64, userID string) (*Transaction, error) {
	userIDUUID, err := uuid.Parse(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(s.Repository.Transactions)
	newTransaction, err := s.Repository.Transactions.Add(transactionType, amount, userIDUUID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &Transaction{
		ID:        newTransaction.ID.String(),
		Amount:    newTransaction.Amount,
		UserID:    newTransaction.UserID.String(),
		CreatedAt: newTransaction.CreatedAt,
		Type:      newTransaction.Type,
	}, nil
}
