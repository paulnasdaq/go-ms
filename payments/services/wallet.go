package services

import (
	"github.com/paulnasdaq/fms-v2/common"
	"github.com/paulnasdaq/fms-v2/payments/db"
)

type WalletService struct {
	db      *db.Repository
	clients *common.ClientManager
}

func NewWalletService(db *db.Repository, clients *common.ClientManager) *WalletService {
	return &WalletService{db, clients}
}

func (s *WalletService) Get(userID string) *Wallet {
	return &Wallet{userID: userID, walletService: s}
}

type Wallet struct {
	userID        string
	walletService *WalletService
}

//func (w *Wallet) Balance() float64 {
//	w.walletService.db.Transactions
//}
