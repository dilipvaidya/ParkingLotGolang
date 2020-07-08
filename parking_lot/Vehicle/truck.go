package vehicle

import (
	ticket "github.com/ParkingLotGolang/parking_lot/ticketservice"
)

// Truck type Vehicle
type Truck struct {
	Vehicle
	vehicleType        Type
	registrationNumber string
	color              string
	ticket             *ticket.Ticket
}

// GetNewVehicle gives Vehicle object of Truck type.
/*func GetNewVehicle(registrationNum string, truckColor string, ticket *ticket.Ticket) Vehicle {
	return &Truck{
		vehicleType:        VehicleTypeTruck,
		registrationNumber: registrationNum,
		color:              truckColor,
		ticket:             ticket,
	}
}*/

// GetVehicleRegistrationNumber returns Truck's reg num
func (t *Truck) GetVehicleRegistrationNumber() string {
	return t.registrationNumber
}

// GetVehicleColor return color of the Truck
func (t *Truck) GetVehicleColor() string {
	return t.color
}

// GetTicketDetails return color of the Truck
func (t *Truck) GetTicketDetails() *ticket.Ticket {
	return t.ticket
}

// GetVehicleType return type of the vehicle
func (t *Truck) GetVehicleType() Type {
	return t.vehicleType
}
