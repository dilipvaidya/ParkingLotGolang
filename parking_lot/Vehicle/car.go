package vehicle

import (
	ticket "../ticketservice"
)

// Car type Vehicle
type Car struct {
	*Vehicle
	vehicleType Type
}

// GetNewCar gives Vehicle object of car type.
func GetNewCar(registrationNum string, carColor string, ticket *ticket.Ticket) *Car {
	var t Car
	t.Vehicle = GetNewVehicle(registrationNum, carColor, ticket)
	t.vehicleType = VehicleTypeCar
	return &t
}

// GetVehicleRegistrationNumber returns car's reg num
func (c *Car) GetVehicleRegistrationNumber() string {
	return c.Vehicle.GetVehicleRegistrationNumber()
}

// GetVehicleColor return color of the Car
func (c *Car) GetVehicleColor() string {
	return c.Vehicle.GetVehicleColor()
}

// GetTicketDetails return color of the Car
func (c *Car) GetTicketDetails() *ticket.Ticket {
	return c.Vehicle.GetTicketDetails()
}
