package commandservice

import (
	"errors"
	"strconv"

	parkingservice "github.com/ParkingLotGolang/parking_lot/parkingservice"
	vehicle "github.com/ParkingLotGolang/parking_lot/vehicle"
)

// CommandPark class
type CommandPark struct {
	Command
	discreption        string
	commandArray       []string
	slotToParkIn       int
	vehicleNumberPlate string
	vehicleColor       string
	vehicleType        vehicle.Type
}

// Park procedure
func Park(createParkingCommand []string) Command {
	return &CommandPark{
		discreption:  "ParkingSlot Creation Command",
		commandArray: createParkingCommand,
	}
}

// GetSlotToParkIn returns slots to be created
func (p *CommandPark) GetSlotToParkIn() int {
	return p.slotToParkIn
}

// SetSlotToParkIn sets slots to be created
func (p *CommandPark) SetSlotToParkIn(numOfSlots int) {
	p.slotToParkIn = numOfSlots
}

// GetCommandArray returns commandArray class var
func (p *CommandPark) GetCommandArray() []string {
	return p.commandArray
}

// SetVehicleRegistrationNumber sets the number plate for vehicle
func (p *CommandPark) SetVehicleRegistrationNumber(regNumber string) {
	p.vehicleNumberPlate = regNumber
}

// GetVehicleRegistrationNumber sets the number plate for vehicle
func (p *CommandPark) GetVehicleRegistrationNumber() string {
	return p.vehicleNumberPlate
}

// SetVehicleColor sets the number plate for vehicle
func (p *CommandPark) SetVehicleColor(color string) {
	p.vehicleColor = color
}

// GetVehicleColor sets the number plate for vehicle
func (p *CommandPark) GetVehicleColor() string {
	return p.vehicleColor
}

// SetVehicleType sets the number plate for vehicle
func (p *CommandPark) SetVehicleType(vType vehicle.Type) {
	p.vehicleType = vType
}

// GetVehicleType sets the number plate for vehicle
func (p *CommandPark) GetVehicleType() vehicle.Type {
	return p.vehicleType
}

// ValidateCommand validates the parking command
func (p *CommandPark) ValidateCommand() error {
	// valide command- $ create_parking_lot <SLOT_NUMBER>
	if p.GetCommandArray() == nil && len(p.GetCommandArray()) < 3 {
		return errors.New("Invalid park command format")
	}

	p.SetVehicleRegistrationNumber(p.GetCommandArray()[1])
	p.SetVehicleColor(p.GetCommandArray()[2])

	//Type is made optional
	if len(p.GetCommandArray()) == 3 {
		p.SetVehicleType(vehicle.VehicleTypeCar)
	} else {
		vType, err := strconv.Atoi(p.GetCommandArray()[4])
		if nil != err {
			return errors.New("Number format exception in park command")
		}

		p.SetVehicleType(vehicle.Type(vType))
	}
	return nil
}

// ExecuteCommand executes the parking command
func (p *CommandPark) ExecuteCommand() error {
	pkService := parkingservice.CreateParking(0) // won't be created as its singleton class
	err := pkService.CheckInNewVehicle(p.GetVehicleRegistrationNumber(), p.GetVehicleColor(), p.GetVehicleType())
	if nil != err {
		return err
	}
	return nil
}
