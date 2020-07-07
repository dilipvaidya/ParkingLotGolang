package commandservice

import (
	"errors"

	parkingservice "github.com/ParkingLotGolang/parking_lot/parkingservice"
)

// CommandStatus class
type CommandStatus struct {
	Command
	discreption  string
	commandArray []string
}

// Status procedure
func Status(createParkingCommand []string) Command {
	return &CommandStatus{
		discreption:  "Parking Status Comnand",
		commandArray: createParkingCommand,
	}
}

// GetCommandArray returns commandArray class var
func (p *CommandStatus) GetCommandArray() []string {
	return p.commandArray
}

// ValidateCommand validates the parking command
func (p *CommandStatus) ValidateCommand() error {
	// valide command- $ create_parking_lot <SLOT_NUMBER>
	if p.GetCommandArray() == nil && len(p.GetCommandArray()) != 1 {
		return errors.New("Invalid Status command format")
	}
	return nil
}

// ExecuteCommand executes the parking command
func (p *CommandStatus) ExecuteCommand() error {
	pkService := parkingservice.CreateParking(0) // won't be created as its singleton class
	pkService.DisplayParkingStatus()
	return nil
}
