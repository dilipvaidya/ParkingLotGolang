package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ParkingLotGolang/parking_lot/commandservice"
)

func main() {
	var scanner *bufio.Scanner

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

	for scanner.Scan() {
		nextCommandString := scanner.Text()
		if strings.Compare(nextCommandString, "exit") == 0 {
			break
		} else {
			commandservice.ExecuteCommand(nextCommandString)
		}
	}
}
