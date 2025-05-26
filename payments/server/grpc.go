package server

import (
	"context"
	"github.com/paulnasdaq/fms-v2/common"
	pb "github.com/paulnasdaq/fms-v2/common/transactions"
	"github.com/paulnasdaq/fms-v2/payments/db"
	"github.com/paulnasdaq/fms-v2/payments/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedTransactionsServiceServer
	service *services.TransactionService
	clients *common.ClientManager
}

func New(clients *common.ClientManager, service *services.TransactionService) *Server {
	return &Server{clients: clients, service: service}
}

func (s *Server) ListenAndServe(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("Fialed to start Payments GRPC Server", err)
		return err
	}
	ss := grpc.NewServer()
	pb.RegisterTransactionsServiceServer(ss, s)
	reflection.Register(ss)
	return ss.Serve(lis)
}

func (s *Server) AddTransaction(ctx context.Context, request *pb.AddTransactionRequest) (*pb.AddTransactionResponse, error) {
	newTransaction, err := s.service.Add(db.TransactionType(request.Type), request.Amount, request.User_ID)
	if err != nil {
		log.Println("failed to add transaction", err)
		return nil, err
	}
	return &pb.AddTransactionResponse{
		ID:        newTransaction.ID,
		User_ID:   newTransaction.UserID,
		CreatedAt: timestamppb.New(newTransaction.CreatedAt),
		Amount:    newTransaction.Amount,
		Type:      string(newTransaction.Type),
	}, nil
}
