package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VehicleModel struct {
	gorm.Model
	ID       uuid.UUID
	Name     string
	Vehicles []Vehicle `gorm:"foreignKey:ModelID"`
}

type VehicleModelDomain struct {
	db *gorm.DB
}

func (d *VehicleModelDomain) Add(name string) (*VehicleModel, error) {
	vehicleModel := VehicleModel{Name: name, ID: uuid.New()}
	if res := d.db.Create(&vehicleModel); res.Error != nil {
		return nil, res.Error
	}
	return &vehicleModel, nil
}

func (d *VehicleModelDomain) Get(ID uuid.UUID) (*VehicleModel, error) {
	vehicleModel := VehicleModel{}
	if res := d.db.First(&vehicleModel, ID); res.Error != nil {
		return nil, res.Error
	}
	return &vehicleModel, nil
}
