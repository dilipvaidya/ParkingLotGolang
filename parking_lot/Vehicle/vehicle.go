package vehicle

import (
	ticket "github.com/ParkingLotGolang/parking_lot/ticketservice"
)

type Type int

const (
	VehicleTypeCar Type = iota
	VehicleTypeTruck
)

// Vehicle parent class
type Vehicle struct {
	RegistrationNumber string
	color              string
	Ticket             *ticket.Ticket
}

// GetNewVehicle gives Vehicle object of car type.
func GetNewVehicle(registrationNum string, carColor string, t *ticket.Ticket) *Vehicle {
	return &Vehicle{
		RegistrationNumber: registrationNum,
		color:              carColor,
		Ticket:             t,
	}
}

// GetVehicleRegistrationNumber returns car's reg num
func (v *Vehicle) GetVehicleRegistrationNumber() string {
	return v.RegistrationNumber
}

// GetVehicleColor return color of the Car
func (v *Vehicle) GetVehicleColor() string {
	return v.color
}

// GetTicketDetails return color of the Car
func (v *Vehicle) GetTicketDetails() *ticket.Ticket {
	return v.Ticket
}
