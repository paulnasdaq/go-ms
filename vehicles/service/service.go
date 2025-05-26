package service

import (
	"github.com/google/uuid"
	"github.com/paulnasdaq/fms-v2/vehicles/db"
	"log"
)

type Service interface {
	AddVehicleModel(name string) (*VehicleModel, error)
	AddVehicle(regNumber string, chassisNumber string, modelId string) (*Vehicle, error)
	//AddControllerModel(name string) ControllerModel
	//AddController(serialNumber string) Controller
	//	AddVehicle()
	//	AddController()
	//	BindControllerToVehicle()
}

type VehiclesService struct {
	repository *db.Repository
}

func (s *VehiclesService) AddVehicleModel(name string) (*VehicleModel, error) {
	log.Println("Adding vehicle model")
	v, err := s.repository.VehicleModels.Add(name)
	if err != nil {
		return nil, err
	}
	return &VehicleModel{ID: v.ID.String(), Name: v.Name}, nil
}

func (s *VehiclesService) AddVehicle(regNumber string, chassisNumber string, modelId string) (*Vehicle, error) {
	modelIdUUID, err := uuid.Parse(modelId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v, err := s.repository.Vehicles.Add(regNumber, chassisNumber, modelIdUUID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	vehicle := Vehicle{ID: v.ID.String(), RegNumber: v.RegNumber, ChassisNumber: v.ChassisNumber}
	vehicleModel, err := s.repository.VehicleModels.Get(modelIdUUID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	vehicle.Model = VehicleModel{Name: vehicleModel.Name, ID: vehicleModel.ID.String()}
	return &vehicle, nil
}

//	func (s *VehiclesService) AddControllerModel(name string) ControllerModel {
//		//TODO implement me
//		panic("implement me")
//	}
//
//	func (s *VehiclesService) AddController(serialNumber string) Controller {
//		//TODO implement me
//		panic("implement me")
//	}
func NewVehiclesService(r *db.Repository) Service {
	return &VehiclesService{repository: r}
}

//
//func (s *VehiclesService) AddModel(s) {
//
//}
