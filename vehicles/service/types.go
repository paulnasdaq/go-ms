package service

type Controller struct {
	ID           string
	SerialNumber string
	Model        ControllerModel
}
type VehicleModel struct {
	ID   string
	Name string
}
type ControllerModel struct {
	ID   string
	Name string
}
