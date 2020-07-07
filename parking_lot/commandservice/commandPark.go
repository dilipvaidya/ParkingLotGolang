package commandservice

import (
	"errors"

	parkingservice "github.com/ParkingLotGolang/parking_lot/parkingservice"
)

// CommandPark class
type CommandPark struct {
	Command
	discreption        string
	commandArray       []string
	slotToParkIn       int
	vehicleNumberPlate string
	vehicleColor       string
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

// ValidateCommand validates the parking command
func (p *CommandPark) ValidateCommand() error {
	// valide command- $ create_parking_lot <SLOT_NUMBER>
	if p.GetCommandArray() == nil && len(p.GetCommandArray()) != 3 {
		return errors.New("Invalid park command format")
	}

	p.SetVehicleRegistrationNumber(p.GetCommandArray()[1])
	p.SetVehicleColor(p.GetCommandArray()[2])
	return nil
}

// ExecuteCommand executes the parking command
func (p *CommandPark) ExecuteCommand() error {
	pkService := parkingservice.CreateParking(0) // won't be created as its singleton class
	err := pkService.CheckInNewVehicle(p.GetVehicleRegistrationNumber(), p.GetVehicleColor())
	if nil != err {
		return errors.New(err.Error())
	}
	return nil
}
