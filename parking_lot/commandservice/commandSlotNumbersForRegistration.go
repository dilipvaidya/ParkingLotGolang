package commandservice

import (
	"errors"
	"fmt"

	parkingservice "github.com/ParkingLotGolang/parking_lot/parkingservice"
)

// CommandSlotForReg class
type CommandSlotForReg struct {
	Command
	discreption  string
	commandArray []string
	regNumber    string
}

// GetCommandArray returns commandArray class var
func (r *CommandSlotForReg) GetCommandArray() []string {
	return r.commandArray
}

// SlotForReg procedure
func SlotForReg(createParkingCommand []string) Command {
	return &CommandSlotForReg{
		discreption:  "Parking Allocattion Command",
		commandArray: createParkingCommand,
	}
}

// GetRegNumber returns slots to be created
func (r *CommandSlotForReg) GetRegNumber() string {
	return r.regNumber
}

// SetRegNumber sets slots to be created
func (r *CommandSlotForReg) SetRegNumber(reg string) {
	r.regNumber = reg
}

// ValidateCommand validates the parking command
func (r *CommandSlotForReg) ValidateCommand() error {
	// valide command- $ create_parking_lot <SLOT_NUMBER>
	if r.GetCommandArray() == nil && len(r.GetCommandArray()) != 2 {
		return errors.New("Invalid Leave command format")
	}

	r.SetRegNumber(r.GetCommandArray()[1])
	return nil
}

// ExecuteCommand executes the parking command
func (r *CommandSlotForReg) ExecuteCommand() error {
	pkService := parkingservice.CreateParking(0) // won't be created as its singleton class
	slotNumber, err := pkService.GetSlotForRegistrationNumber(r.GetRegNumber())
	if nil != err {
		return err
	}
	fmt.Println(slotNumber)
	return nil
}
