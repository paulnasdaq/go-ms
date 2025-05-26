package service

import "gorm.io/gorm"

type Vehicle struct {
	ID            string
	RegNumber     string
	ChassisNumber string
	Model         VehicleModel
}

type VehiclesDomain struct {
	db gorm.DB
}

//func (v *Vehicle) Add(registrationNumber string) (*Vehicle, error) {
//
//}
