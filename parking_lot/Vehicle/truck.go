package vehicle

import (
	ticket "github.com/ParkingLotGolang/parking_lot/ticketservice"
)

// Truck type Vehicle
type Truck struct {
	*Vehicle
	vehicleType Type
}

// GetNewTruck gives Vehicle object of Truck type.
func GetNewTruck(registrationNum string, TruckColor string, ticket *ticket.Ticket) *Truck {
	var t Truck
	t.Vehicle = GetNewVehicle(registrationNum, TruckColor, ticket)
	t.vehicleType = VehicleTypeTruck
	return &t
}

// GetVehicleRegistrationNumber returns Truck's reg num
func (t *Truck) GetVehicleRegistrationNumber() string {
	return t.Vehicle.GetVehicleRegistrationNumber()
}

// GetVehicleColor return color of the Truck
func (t *Truck) GetVehicleColor() string {
	return t.Vehicle.GetVehicleColor()
}

// GetTicketDetails return color of the Truck
func (t *Truck) GetTicketDetails() *ticket.Ticket {
	return t.Vehicle.GetTicketDetails()
}
