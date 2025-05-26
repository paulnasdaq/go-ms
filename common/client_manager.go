package common

import (
	authPb "github.com/paulnasdaq/fms-v2/common/auth"
	transactionsPb "github.com/paulnasdaq/fms-v2/common/transactions"
	usersPb "github.com/paulnasdaq/fms-v2/common/users"
	vehiclesPb "github.com/paulnasdaq/fms-v2/common/vehicles"
	"google.golang.org/grpc"
	"log"
)

type ClientManager struct {
	usersClient        usersPb.UsersServiceClient
	TransactionsClient transactionsPb.TransactionsServiceClient
	vehicles           vehiclesPb.VehiclesServiceClient
	auth               authPb.AuthClient
}

func NewClientManager() *ClientManager {
	return &ClientManager{}
}

func (m *ClientManager) Auth() (authPb.AuthClient, error) {
	if m.auth == nil {
		log.Println("Attempting to connect to auth servie")
		con, err := grpc.NewClient(":3003", grpc.WithInsecure())
		if err != nil {
			log.Println("Failed to connect to auth service", err)
			return nil, err
		}
		m.auth = authPb.NewAuthClient(con)
	}
	return m.auth, nil
}
func (m *ClientManager) Users() (usersPb.UsersServiceClient, error) {
	if m.usersClient == nil {
		log.Println("Attempting to connect to users service")
		con, err := grpc.NewClient(":3000", grpc.WithInsecure())
		if err != nil {
			log.Println("Failed to connect to users service", err)
			return nil, err
		}

		m.usersClient = usersPb.NewUsersServiceClient(con)
	}
	return m.usersClient, nil
}

func (m *ClientManager) GetTransactionsClient() (transactionsPb.TransactionsServiceClient, error) {
	if m.TransactionsClient == nil {
		log.Println("Attempting to connect to the transactions service")
		con, err := grpc.NewClient(":3001", grpc.WithInsecure())
		if err != nil {
			log.Println("Failed to connect to the transactions service")
			return nil, err
		}
		m.TransactionsClient = transactionsPb.NewTransactionsServiceClient(con)
	}
	return m.TransactionsClient, nil
}

func (m *ClientManager) GetVehiclesClient() (vehiclesPb.VehiclesServiceClient, error) {
	if m.vehicles == nil {
		log.Println("Attempting to connect to the vehicles microservice")
		con, err := grpc.NewClient(":3002", grpc.WithInsecure())
		if err != nil {
			log.Println("Failed to connect to the vehicles service")
			return nil, err
		}
		m.vehicles = vehiclesPb.NewVehiclesServiceClient(con)
	}
	return m.vehicles, nil
}
