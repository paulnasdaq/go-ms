package main

import (
	"github.com/paulnasdaq/fms-v2/common"
	"github.com/paulnasdaq/fms-v2/payments/db"
	"github.com/paulnasdaq/fms-v2/payments/server"
	"github.com/paulnasdaq/fms-v2/payments/services"
	"log"
)

func main() {
	repo, err := db.NewRepository()
	if err != nil {
		log.Printf("Failed to initialize repo", err)
		return
	}

	log.Println(repo.Transactions)
	s := services.TransactionService{Repository: repo}
	client := common.NewClientManager()
	grpcServer := server.New(client, &s)

	grpcServer.ListenAndServe(":3001")
}
