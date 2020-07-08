package commandservice

import (
	"errors"
	"fmt"

	parkingservice "github.com/ParkingLotGolang/parking_lot/parkingservice"
)

// CommandSlotWithVehicleColor class
type CommandSlotWithVehicleColor struct {
	Command
	discreption  string
	commandArray []string
	color        string
}

// GetCommandArray returns commandArray class var
func (r *CommandSlotWithVehicleColor) GetCommandArray() []string {
	return r.commandArray
}

// SlotWithVehicleColor procedure
func SlotWithVehicleColor(createParkingCommand []string) Command {
	return &CommandSlotWithVehicleColor{
		discreption:  "Parking Allocattion Command",
		commandArray: createParkingCommand,
	}
}

// GetColor returns slots to be created
func (r *CommandSlotWithVehicleColor) GetColor() string {
	return r.color
}

// SetColor sets slots to be created
func (r *CommandSlotWithVehicleColor) SetColor(color string) {
	r.color = color
}

// ValidateCommand validates the parking command
func (r *CommandSlotWithVehicleColor) ValidateCommand() error {
	// valide command- $ create_parking_lot <SLOT_NUMBER>
	if r.GetCommandArray() == nil && len(r.GetCommandArray()) != 2 {
		return errors.New("Invalid CarsWithColor command format")
	}

	r.SetColor(r.GetCommandArray()[1])
	return nil
}

// ExecuteCommand executes the parking command
func (r *CommandSlotWithVehicleColor) ExecuteCommand() error {
	pkService := parkingservice.CreateParking(0) // won't be created as its singleton class
	slots, err := pkService.GetSlotForVehicleColor(r.GetColor())
	if nil != err {
		return err
	}

	for index, element := range slots {
		if index == len(slots)-1 {
			fmt.Printf("%d", element)
		} else {
			fmt.Printf("%d, ", element)
		}
	}
	fmt.Println()
	return nil
}
