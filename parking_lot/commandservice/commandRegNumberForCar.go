package commandservice

import (
	"errors"
	"fmt"
	"strings"

	parkingservice "github.com/ParkingLotGolang/parking_lot/parkingservice"
)

// CommandCarsWithColor class
type CommandCarsWithColor struct {
	Command
	discreption  string
	commandArray []string
	color        string
}

// GetCommandArray returns commandArray class var
func (r *CommandCarsWithColor) GetCommandArray() []string {
	return r.commandArray
}

// CarsWithColor procedure
func CarsWithColor(createParkingCommand []string) Command {
	return &CommandCarsWithColor{
		discreption:  "Parking Allocattion Command",
		commandArray: createParkingCommand,
	}
}

// GetColor returns slots to be created
func (r *CommandCarsWithColor) GetColor() string {
	return r.color
}

// SetColor sets slots to be created
func (r *CommandCarsWithColor) SetColor(color string) {
	r.color = color
}

// ValidateCommand validates the parking command
func (r *CommandCarsWithColor) ValidateCommand() error {
	// valide command- $ create_parking_lot <SLOT_NUMBER>
	if r.GetCommandArray() == nil && len(r.GetCommandArray()) != 2 {
		return errors.New("Invalid CarsWithColor command format")
	}

	r.SetColor(r.GetCommandArray()[1])
	return nil
}

// ExecuteCommand executes the parking command
func (r *CommandCarsWithColor) ExecuteCommand() error {
	pkService := parkingservice.CreateParking(0) // won't be created as its singleton class
	regStrings, err := pkService.GetCarsRegistrationNumberWithColor(r.GetColor())
	if nil != err {
		return err
	}
	fmt.Printf("%s", strings.Join(regStrings, ", "))
	fmt.Println()
	return nil
}
