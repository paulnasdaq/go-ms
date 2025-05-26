package server

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/paulnasdaq/fms-v2/common"
	trPb "github.com/paulnasdaq/fms-v2/common/transactions"
	"log"
	"net/http"
	"time"
)

type Transaction struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userID"`
	Amount    float64   `json:"amount"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
}

func GetHandler(clients *common.ClientManager) (*chi.Mux, error) {
	router := chi.NewRouter()

	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.Logger)

	router.Post("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		requestBody := Transaction{}
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Printf("Could not decode request", err)
			http.Error(responseWriter, "Could not ready request", http.StatusBadRequest)
			return
		}
		transactionsClient, err := clients.GetTransactionsClient()
		if err != nil {
			log.Println("Could not get transactions serviece", err)
			http.Error(responseWriter, "Service unavailable", http.StatusBadGateway)
			return
		}
		ctx := context.Context(context.Background())
		res, err := transactionsClient.AddTransaction(ctx, &trPb.AddTransactionRequest{Type: requestBody.Type, Amount: requestBody.Amount, User_ID: requestBody.UserID})
		if err != nil {
			log.Println("Could not create transactions serviece", err)
			http.Error(responseWriter, "Service unavailable", http.StatusBadGateway)
			return
		}
		if err = json.NewEncoder(responseWriter).Encode(&Transaction{
			ID:        res.ID,
			UserID:    res.User_ID,
			Amount:    res.Amount,
			CreatedAt: res.CreatedAt.AsTime(),
			Type:      res.Type,
		}); err != nil {
			log.Println("Could not get transactions serviece", err)
			http.Error(responseWriter, "Service unavailable", http.StatusInternalServerError)
			return
		}

	})
	return router, nil

}
