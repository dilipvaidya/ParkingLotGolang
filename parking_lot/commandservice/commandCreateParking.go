package commandservice

import (
	"errors"
	"strconv"

	parkingservice "github.com/ParkingLotGolang/parking_lot/parkingservice"
)

// CommandParking class
type CommandParking struct {
	Command
	discreption   string
	commandArray  []string
	slotsToCreate int
}

// CreateParking creates parking
func CreateParking(createParkingCommand []string) Command {
	return &CommandParking{
		discreption:  "Parking Creation Command",
		commandArray: createParkingCommand,
	}
}

// GetSlotsToCreate returns slots to be created
func (p *CommandParking) GetSlotsToCreate() int {
	return p.slotsToCreate
}

// SetSlotsToCreate sets slots to be created
func (p *CommandParking) SetSlotsToCreate(numOfSlots int) {
	p.slotsToCreate = numOfSlots
}

// GetCommandArray returns commandArray class var
func (p *CommandParking) GetCommandArray() []string {
	return p.commandArray
}

// ValidateCommand validates the parking command
func (p *CommandParking) ValidateCommand() error {
	// valide command- $ create_parking_lot <SLOT_NUMBER>
	if p.GetCommandArray() == nil && len(p.GetCommandArray()) != 2 {
		return errors.New("Invalid creation_parking_lot format")
	}

	slots, err := strconv.Atoi(p.GetCommandArray()[1])
	if nil != err {
		return errors.New("Number format exception in create_parking_lot command")
	}

	p.SetSlotsToCreate(slots)
	return nil
}

// ExecuteCommand executes the parking command
func (p *CommandParking) ExecuteCommand() error {
	parkingservice.CreateParking(p.GetSlotsToCreate())
	return nil
}
