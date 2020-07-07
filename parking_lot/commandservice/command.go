package commandservice

import (
	"fmt"
	"strings"

	utils "github.com/ParkingLotGolang/parking_lot/utils"
)

const (
	err001 string = "Invalid command:'%s'\n"
	err002 string = "Command not supported:'%s'\n"
)

// Command interface
type Command interface {
	GetParkingService()
	ValidateCommand() error
	ExecuteCommand() error
}

func getCommand(commandLine []string) Command {

	switch commandLine[0] {
	case "create_parking_lot":
		return CreateParking(commandLine)

	case "park":
		return Park(commandLine)

	case "leave":
		return Leave(commandLine)

	case "status":
		return Status(commandLine)

	case "registration_numbers_for_cars_with_colour":
	case "slot_numbers_for_cars_with_colour":
	case "slot_number_for_registration_number":
	default:
		fmt.Printf(err002, commandLine[0])
		return nil
	}

	return nil
}

// ExecuteCommand executes each command for parking lot
func ExecuteCommand(commandStr string) bool {
	command := ""
	if command = strings.Trim(commandStr, command); len(command) == 0 {
		fmt.Printf(err001, command)
		return false
	}

	commandLine := strings.Split(command, utils.SPACE)
	commandToExecute := getCommand(commandLine)
	if nil == commandToExecute {
		fmt.Printf(err002, command)
		return false
	}

	// validate the command
	if err := commandToExecute.ValidateCommand(); nil != err {
		fmt.Println(err.Error())
		return false
	}

	// execute command
	err := commandToExecute.ExecuteCommand()
	if nil != err {
		fmt.Println(err.Error())
		return false
	}

	return true
}
