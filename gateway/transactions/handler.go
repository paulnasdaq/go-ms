package transactions

import (
	//"github.com/paulnasdaq/fms-v2/common"
	"github.com/paulnasdaq/fms-v2/payments/db"
	"time"
)

type Transaction struct {
	ID        string  `json:"id"`
	UserID    string  `json:"userID"`
	Amount    float64 `json:"amount"`
	Type      db.TransactionType
	CreatedAt time.Time `json:"createdAt"`
}

//func TransactionsRrequestHandler() (*chi.Mux, error) {
//	//clients := common.ClientManager{}
//	router := chi.NewRouter()
//
//	//s, err := clients.GetTransactionsClient()
//
//	router.Use(middleware.AllowContentType("application/json"))
//	router.Use(middleware.Logger)
//
//	router.Post("/transactions", func(responseWriter http.ResponseWriter, request *http.Request) {
//
//	})
//}
