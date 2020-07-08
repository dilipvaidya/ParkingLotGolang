package commandservice

import (
	"errors"
	"strconv"

	parkingservice "github.com/ParkingLotGolang/parking_lot/parkingservice"
)

// CommandLeave class
type CommandLeave struct {
	Command
	discreption    string
	commandArray   []string
	slotToCheckOut int
}

// Leave procedure
func Leave(createParkingCommand []string) Command {
	return &CommandLeave{
		discreption:  "Parking Allocattion Command",
		commandArray: createParkingCommand,
	}
}

// GetSlotToCheckOut returns slots to be created
func (p *CommandLeave) GetSlotToCheckOut() int {
	return p.slotToCheckOut
}

// SetSlotToCheckOut sets slots to be created
func (p *CommandLeave) SetSlotToCheckOut(slt int) {
	p.slotToCheckOut = slt
}

// GetCommandArray returns commandArray class var
func (p *CommandLeave) GetCommandArray() []string {
	return p.commandArray
}

// ValidateCommand validates the parking command
func (p *CommandLeave) ValidateCommand() error {
	// valide command- $ create_parking_lot <SLOT_NUMBER>
	if p.GetCommandArray() == nil && len(p.GetCommandArray()) != 2 {
		return errors.New("Invalid Leave command format")
	}

	slot, err := strconv.Atoi(p.GetCommandArray()[1])
	if nil != err {
		return errors.New("Number format exception in create_parking_lot command")
	}

	p.SetSlotToCheckOut(slot)

	return nil
}

// ExecuteCommand executes the parking command
func (p *CommandLeave) ExecuteCommand() error {
	pkService := parkingservice.CreateParking(0) // won't be created as its singleton class
	err := pkService.CheckOutVehicle(p.GetSlotToCheckOut())
	if nil != err {
		return err
	}
	return nil
}
