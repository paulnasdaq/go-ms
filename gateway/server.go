package main

import (
	"context"
	"encoding/json"
	"github.com/paulnasdaq/fms-v2/common"
	pb "github.com/paulnasdaq/fms-v2/common/auth"
	"log"
	"net/http"
)

type Server struct {
	clients *common.ClientManager
}

func NewServer() *Server {
	clients := common.NewClientManager()
	return &Server{clients: clients}
}

type User_ struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

func (s *Server) RegisterRoutes() {
	http.HandleFunc("/api/v1/auth/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var user User_
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, "Could not read body", http.StatusBadRequest)
				return
			}
			authClient, err := s.clients.Auth()
			if err != nil {
				log.Println("Failed to get auth client", err)
				http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
				return
			}

			newUser, err := authClient.AddUser(context.Context(context.Background()), &pb.AddUserRequest{Email: user.Email, Password: user.Password})
			if err != nil {
				log.Println("Failed to create user", err)
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}
			if err = json.NewEncoder(w).Encode(&User_{ID: newUser.Id, Email: newUser.Email}); err != nil {
				log.Println("Failed to return")
				http.Error(w, "Failed again", 500)
			}
		}

	})

	http.HandleFunc("/api/v1/auth/tokens", func(w http.ResponseWriter, r *http.Request) {
		var user User_
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Println("Failed to decode JSON", err)
			http.Error(w, "Failed to load json", 400)
			return
		}
		authClient, err := s.clients.Auth()
		if err != nil {
			log.Println("Failed to get auth service", err)
			http.Error(w, "Service unavailable", 504)
			return
		}

		res, err := authClient.GetToken(context.Context(context.Background()), &pb.GetTokenRequest{Email: user.Email, Password: user.Password})
		if err != nil {
			log.Println("Failed to get token", err)
			http.Error(w, "Invalid login creds", 401)
			return
		}
		if err = json.NewEncoder(w).Encode(&Token{Token: res.Token}); err != nil {
			log.Println("failed to return", err)
			http.Error(w, "Stuff", 500)
		}
	})
}
