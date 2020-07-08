package parkingservice

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	slot "github.com/ParkingLotGolang/parking_lot/parkingslot"
	ticket "github.com/ParkingLotGolang/parking_lot/ticketservice"
	utils "github.com/ParkingLotGolang/parking_lot/utils"
	"github.com/ParkingLotGolang/parking_lot/vehicle"
)

// SlotAllocationError enum holds all the slot allocation de-allocation related errors
type SlotAllocationError string

const (
	errparkingfull      string = "Sorry, parking lot is full"
	errslotnotfound     string = "Parking slot is either already empty or not in use"
	errslotdoesnotexist string = "Parking slot does exist"
)

var once sync.Once
var parking *ParkingService

// TypeSlotAllocationMap is a map of slot to its allocation status in the Parking
type TypeSlotAllocationMap map[int]*slot.Slot

// TypeSlotVehicleAllocationMap maps vehicle license number to ticket issued
type TypeSlotVehicleAllocationMap map[int]vehicle.Vehicle

// TypeCheckedInVehiclesMap maintains the map of vehicle to slot number - reverse map
type TypeCheckedInVehiclesMap map[string]int

// ParkingService is the service to create parking and allocate one to vehicle
type ParkingService struct {
	numberOfSlots           int
	slotAllocation          TypeSlotAllocationMap
	slotVehicleAllocation   TypeSlotVehicleAllocationMap
	vehicleToSlotAllocation TypeCheckedInVehiclesMap
}

// CreateParking singletone method to create parking with slots.
func CreateParking(totalNumberOfSlots int) *ParkingService {

	once.Do(func() {
		parking = &ParkingService{
			numberOfSlots:           totalNumberOfSlots + 1,
			slotAllocation:          make(TypeSlotAllocationMap, totalNumberOfSlots+1),
			slotVehicleAllocation:   make(TypeSlotVehicleAllocationMap, totalNumberOfSlots+1),
			vehicleToSlotAllocation: make(TypeCheckedInVehiclesMap, totalNumberOfSlots+1),
		}

		for i := 1; i <= totalNumberOfSlots; i++ {
			parking.slotAllocation[i] = slot.GetNewSlot(i, slot.Free)
		}
	})
	return parking
}

// GetNumberOfSlots return the number of slots in the Parking
func (p *ParkingService) GetNumberOfSlots() int {
	return p.numberOfSlots
}

// GetSlotAllocationMap returns slot allocation map
func (p *ParkingService) GetSlotAllocationMap() TypeSlotAllocationMap {
	return p.slotAllocation
}

// CheckInNewVehicle checks in new vehicle, done slot reservation and ticket generation.
func (p *ParkingService) CheckInNewVehicle(licenseNumber string, color string, vehicleType vehicle.Type) error {

	if _, ok := p.vehicleToSlotAllocation[licenseNumber]; ok {
		return errors.New("Vehicle already checked in")
	}

	currentTimeInMilli := utils.CurrentTimeInMilli()
	t := ticket.GetNewTicket(licenseNumber, currentTimeInMilli, 0, currentTimeInMilli, 1.1)
	v := vehicle.GetNewVehicle(licenseNumber, color, t, vehicleType)

	sltNum, err := p.GetNextFreeParkingSlot()
	if nil != err {
		return err
	}

	// allocate slot for current vehicle

	p.slotVehicleAllocation[sltNum] = v
	p.vehicleToSlotAllocation[licenseNumber] = sltNum
	fmt.Println("Allocated slot number: ", sltNum)
	return nil
}

// CheckOutVehicle checks in new vehicle, done slot reservation and ticket generation.
func (p *ParkingService) CheckOutVehicle(sltNum int) error {

	err := p.FreeParkingSlot(sltNum)
	if nil != err {
		return err
	}

	if v, ok := p.slotVehicleAllocation[sltNum]; ok {
		delete(p.vehicleToSlotAllocation, v.GetVehicleRegistrationNumber())
	}

	delete(p.slotVehicleAllocation, sltNum)
	fmt.Printf("Slot number %d is free\n", sltNum)
	return nil
}

// GetNextFreeParkingSlot return the Next Free Slot to be allocated by the car
func (p *ParkingService) GetNextFreeParkingSlot() (int, error) {
	for i := 1; i < p.GetNumberOfSlots(); i++ {
		if p.GetSlotAllocationMap()[i].GetSlotStatus() == slot.Free {
			p.GetSlotAllocationMap()[i].SetSlotStatus(slot.Occupied)
			return i, nil
		}
	}
	return 0, errors.New(errparkingfull)
}

// FreeParkingSlot free's the slot for next time use.
func (p *ParkingService) FreeParkingSlot(slotNumber int) error {
	if slotBeingFreed, ok := p.GetSlotAllocationMap()[slotNumber]; ok {
		if slotBeingFreed.GetSlotStatus() == slot.Free || slotBeingFreed.GetSlotStatus() == slot.NotInUsed {
			return errors.New(errslotnotfound)
		}
		if slotBeingFreed.GetSlotStatus() == slot.Occupied {
			slotBeingFreed.SetSlotStatus(slot.Free)
			return nil
		}
	}
	return errors.New(errslotdoesnotexist)
}

// DisplayParkingStatus displays parking status on stdio so be carefull
func (p *ParkingService) DisplayParkingStatus() {
	fmt.Printf("Slot No.\tRegistration No\tColour\n")
	for i := 1; i < p.GetNumberOfSlots(); i++ {
		if vehicle, ok := p.slotVehicleAllocation[i]; ok {
			fmt.Printf("    %d\t\t%s\t%s\n", i, vehicle.GetVehicleRegistrationNumber(), vehicle.GetVehicleColor())
		}
	}
}

// GetSlotForVehicleColor returns all the cars with specific color
func (p *ParkingService) GetSlotForVehicleColor(color string) ([]int, error) {
	var slots []int

	for slt, v := range p.slotVehicleAllocation {
		if p.GetSlotAllocationMap()[slt].GetSlotStatus() == slot.Occupied &&
			strings.Compare(v.GetVehicleColor(), color) == 0 {
			slots = append(slots, slt)
		}
	}

	if len(slots) == 0 {
		return nil, errors.New("Not found")
	}
	return slots, nil
}

// GetCarsRegistrationNumberWithColor return all cars registration number with specified color
func (p *ParkingService) GetCarsRegistrationNumberWithColor(color string) ([]string, error) {
	var regStr []string

	for slt, v := range p.slotVehicleAllocation {
		if p.GetSlotAllocationMap()[slt].GetSlotStatus() == slot.Occupied &&
			strings.Compare(v.GetVehicleColor(), color) == 0 {
			regStr = append(regStr, v.GetVehicleRegistrationNumber())
		}
	}

	if len(regStr) == 0 {
		return nil, errors.New("Not found")
	}
	return regStr, nil
}

// GetSlotForRegistrationNumber returns slot number for specified vehicle registration number
func (p *ParkingService) GetSlotForRegistrationNumber(regNumber string) (int, error) {

	if slotNumber, ok := p.vehicleToSlotAllocation[regNumber]; ok {
		return slotNumber, nil
	}
	return -1, errors.New("Not found")
}
