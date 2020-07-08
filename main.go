package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ParkingLotGolang/parking_lot/commandservice"
)

const (
	exitCommand = "exit"
)

func main() {
	var scanner *bufio.Scanner

	// get scanner to either scan commands from command line or from input file.
	switch len(os.Args) {

	case 1: //set scanner to read from stdio
		scanner = bufio.NewScanner(os.Stdin)

	case 2: //set scanner to read from file
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)

	default: //error
		scanner = nil
	}

	// inifinite loop until exit command
	for scanner.Scan() {
		nextCommandString := scanner.Text()
		if strings.Compare(nextCommandString, exitCommand) == 0 {
			break
		} else {
			commandservice.ExecuteCommand(nextCommandString)
		}
	}
}
