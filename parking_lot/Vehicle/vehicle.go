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
type Vehicle interface {
	GetVehicleRegistrationNumber() string
	GetVehicleColor() string
	GetTicketDetails() *ticket.Ticket
	GetVehicleType() Type
}

// GetNewVehicle gives Vehicle object of car type.
func GetNewVehicle(registrationNum string, color string, ticket *ticket.Ticket, vehicleType Type) Vehicle {

	switch vehicleType {
	case VehicleTypeCar:
		return &Car{
			vehicleType:        VehicleTypeCar,
			registrationNumber: registrationNum,
			color:              color,
			ticket:             ticket,
		}

	case VehicleTypeTruck:
		return &Truck{
			vehicleType:        VehicleTypeTruck,
			registrationNumber: registrationNum,
			color:              color,
			ticket:             ticket,
		}
	}

	// default add vehicle as of type car
	return &Car{
		vehicleType:        VehicleTypeCar,
		registrationNumber: registrationNum,
		color:              color,
		ticket:             ticket,
	}
}
