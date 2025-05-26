package vehicles

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	pb "github.com/paulnasdaq/fms-v2/common/vehicles"
	"net/http"
)

type VehicleModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Vehicle struct {
	ID                 string       `json:"id"`
	RegistrationNumber string       `json:"registrationNumber"`
	ChassisNumber      string       `json:"chassisNumber"`
	ModelID            string       `json:"modelID"`
	Model              VehicleModel `json:"model"`
}

func GetVehicleModelRoutes(client pb.VehiclesServiceClient) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Logger)

	r.Post("/models", func(w http.ResponseWriter, r *http.Request) {
		var req VehicleModel
		err := json.NewDecoder(r.Body).Decode(&req)
		fmt.Println(req)
		if err != nil {
			http.Error(w, "Could not decode json", http.StatusBadRequest)
			return
		}
		ctx := context.Background()
		model, err := client.AddVehicleModel(ctx, &pb.AddVehicleModelRequest{Name: req.Name})

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Could not decode json", http.StatusBadRequest)
			return
		}
		if err = json.NewEncoder(w).Encode(&VehicleModel{
			Name: model.Name,
			ID:   model.Id,
		}); err != nil {
			http.Error(w, "Could not decode json", http.StatusBadRequest)
			return
		}
	})

	return r
}
func GetVehiclesRoutes(client pb.VehiclesServiceClient) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Logger)

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var requestBody Vehicle
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Could not read request body", http.StatusBadRequest)
			return
		}
		ctx := context.Background()
		vehicle, err := client.AddVehicle(ctx, &pb.AddVehicleRequest{RegistrationNumber: requestBody.RegistrationNumber, ModelId: requestBody.ModelID, ChassisNumber: requestBody.ChassisNumber})
		if err != nil {
			http.Error(w, "Could not create vehicle", http.StatusInternalServerError)
			return
		}
		if err = json.NewEncoder(w).Encode(&Vehicle{
			ID:                 vehicle.Id,
			ChassisNumber:      vehicle.ChassisNumber,
			RegistrationNumber: vehicle.RegistrationNumber,
			ModelID:            vehicle.ModelId,
		}); err != nil {

		}
	})

	return r
}
