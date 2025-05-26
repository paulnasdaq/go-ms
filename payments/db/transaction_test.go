package db

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestTransactionRepository_Balance(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&Transaction{})

	userID := uuid.New()
	t.Run("Test balance", func(t *testing.T) {
		testTransactionRepo := TransactionRepository{db: db}
		userTransactions := []Transaction{
			{Type: Credit, Amount: 100, UserID: userID, ID: uuid.New()},
			{Type: Credit, Amount: 100, UserID: userID, ID: uuid.New()},
			{Type: Debit, Amount: 100, UserID: userID, ID: uuid.New()},
			{Type: Credit, Amount: 100, UserID: userID, ID: uuid.New()},
		}
		db.Create(&userTransactions)
		balance, _ := testTransactionRepo.Balance(uuid.New())
		assert.Equal(t, 200.0, balance)
	})
}
