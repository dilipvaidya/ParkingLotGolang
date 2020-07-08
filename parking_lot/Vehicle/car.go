package vehicle

import (
	ticket "github.com/ParkingLotGolang/parking_lot/ticketservice"
)

// Car type Vehicle
type Car struct {
	Vehicle
	vehicleType        Type
	registrationNumber string
	color              string
	ticket             *ticket.Ticket
}

// GetNew gives Vehicle object of car type.
/*func GetNewVehicle(registrationNum string, carColor string, ticket *ticket.Ticket) Vehicle {
	return &Car{
		vehicleType:        VehicleTypeCar,
		registrationNumber: registrationNum,
		color:              carColor,
		ticket:             ticket,
	}
}*/

// GetVehicleRegistrationNumber returns car's reg num
func (c *Car) GetVehicleRegistrationNumber() string {
	return c.registrationNumber
}

// GetVehicleColor return color of the Car
func (c *Car) GetVehicleColor() string {
	return c.color
}

// GetTicketDetails return color of the Car
func (c *Car) GetTicketDetails() *ticket.Ticket {
	return c.Vehicle.GetTicketDetails()
}

// GetVehicleType return type of the vehicle
func (c *Car) GetVehicleType() Type {
	return c.vehicleType
}
