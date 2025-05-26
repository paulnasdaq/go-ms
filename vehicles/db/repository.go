package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//type baseModel struct {
//	gorm.Model
//	ID uuid.UUID
//}

//type Repository interface {
//	AddVehicleModel(name string) (*VehicleModel, error)
//	AddVehicle(regNumber string, chassisNumber string, modelID string) (*Vehicle, error)
//	GetVehicleModel(id string) (*VehicleModel, error)
//}

type Repository struct {
	db            *gorm.DB
	Vehicles      VehiclesDomain
	VehicleModels VehicleModelDomain
}

func NewRepository() (*Repository, error) {
	db, err := gorm.Open(sqlite.Open("vehicles.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&Controller{}); err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&ControllerModel{}); err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&Vehicle{}); err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&VehicleModel{}); err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&VehicleController{}); err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&VehicleUser{}); err != nil {
		return nil, err
	}
	return &Repository{db: db, Vehicles: VehiclesDomain{db: db}, VehicleModels: VehicleModelDomain{db: db}}, nil
}

//func (r *Repository) AddVehicleModel(name string) (*VehicleModel, error) {
//	vehicleModel := VehicleModel{Name: name}
//	vehicleModel.ID = uuid.New()
//
//	res := r.db.Create(&vehicleModel)
//	if res.Error != nil {
//		log.Println(res.Error)
//		return nil, res.Error
//	}
//	return &vehicleModel, nil
//}
//func (r *Repository) AddVehicle(regNumber string, chassisNumber string, modelID string) (*Vehicle, error) {
//	modelIDUUID, err := uuid.Parse(modelID)
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//	vehicle := Vehicle{ID: uuid.New(), RegNumber: regNumber, ChassisNumber: chassisNumber, ModelID: modelIDUUID}
//	res := r.db.Create(&vehicle)
//	if res.Error != nil {
//		log.Println(res.Error)
//		return nil, res.Error
//	}
//	return &vehicle, nil
//}
//
//func (r *Repository) GetVehicleModel(id string) (*VehicleModel, error) {
//	v := VehicleModel{}
//	res := r.db.First(&v, id)
//	if res.Error != nil {
//		log.Println(res.Error)
//		return nil, res.Error
//	}
//	return &v, nil
//}
