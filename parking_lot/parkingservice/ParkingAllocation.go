package parkingservice

import (
	"errors"
	"fmt"
	"sync"

	slot "github.com/ParkingLotGolang/parking_lot/parkingslot"
	ticket "github.com/ParkingLotGolang/parking_lot/ticketservice"
	utils "github.com/ParkingLotGolang/parking_lot/utils"
	"github.com/ParkingLotGolang/parking_lot/vehicle"
)

// SlotAllocationError enum holds all the slot allocation de-allocation related errors
type SlotAllocationError string

const (
	errparkingfull  string = "Sorry, parking lot is full"
	errslotnotfound string = "Parking slot is either already empty or not in use"
)

var once sync.Once
var parking *ParkingService

// TypeSlotAllocationMap is a map of slot to its allocation status in the Parking
type TypeSlotAllocationMap map[int]*slot.Slot

// TypeSlotVehicleAllocationMap maps vehicle license number to ticket issued
type TypeSlotVehicleAllocationMap map[int]*vehicle.Vehicle

// TypeCheckedInVehiclesMap maintains the map of vehicle to slot numbee - reverse map
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
			numberOfSlots:           totalNumberOfSlots,
			slotAllocation:          make(TypeSlotAllocationMap, totalNumberOfSlots),
			slotVehicleAllocation:   make(TypeSlotVehicleAllocationMap, totalNumberOfSlots),
			vehicleToSlotAllocation: make(TypeCheckedInVehiclesMap, totalNumberOfSlots),
		}

		for i := 0; i < totalNumberOfSlots; i++ {
			parking.slotAllocation[i] = slot.GetNewSlot(i, slot.Free)
		}
	})
	return parking
}

// CheckInNewVehicle checks in new vehicle, done slot reservation and ticket generation.
func (p *ParkingService) CheckInNewVehicle(licenseNumber string, color string) error {

	if _, ok := p.vehicleToSlotAllocation[licenseNumber]; ok {
		return errors.New("Vehicle already checked in")
	}

	currentTimeInMilli := utils.CurrentTimeInMilli()
	t := ticket.GetNewTicket(licenseNumber, currentTimeInMilli, 0, currentTimeInMilli, 1.1)
	v := vehicle.GetNewVehicle(licenseNumber, color, t)

	sltNum, err := p.GetNextFreeParkingSlot()
	if nil != err {
		return err
	}

	// allocate slot for current vehicle

	p.slotVehicleAllocation[sltNum] = v
	p.vehicleToSlotAllocation[licenseNumber] = sltNum
	return nil
}

// CheckOutVehicle checks in new vehicle, done slot reservation and ticket generation.
func (p *ParkingService) CheckOutVehicle(sltNum int) error {

	if v, ok := p.slotVehicleAllocation[sltNum]; ok {
		delete(p.vehicleToSlotAllocation, v.GetVehicleRegistrationNumber())
	}

	delete(p.slotVehicleAllocation, sltNum)
	p.FreeParkingSlot(sltNum)

	return nil
}

// GetNumberOfSlots return the number of slots in the Parking
func (p *ParkingService) GetNumberOfSlots() int {
	return p.numberOfSlots
}

// GetSlotAllocationMap returns slot allocation map
func (p *ParkingService) GetSlotAllocationMap() TypeSlotAllocationMap {
	return p.slotAllocation
}

// GetNextFreeParkingSlot return the Next Free Slot to be allocated by the car
func (p *ParkingService) GetNextFreeParkingSlot() (int, error) {
	for i := 0; i < p.GetNumberOfSlots(); i++ {
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
		if slotBeingFreed.GetSlotStatus() == slot.Occupied {
			slotBeingFreed.SetSlotStatus(slot.Free)
			return nil
		}
	}
	return errors.New(errslotnotfound)
}

// DisplayParkingStatus displays parking status on stdio so be carefull
func (p *ParkingService) DisplayParkingStatus() {
	fmt.Printf("Slot No.\tRegistration No\tColour\n")
	for slot, vehicle := range p.slotVehicleAllocation {
		fmt.Printf("    %d\t\t%s\t%s\n", slot, vehicle.GetVehicleRegistrationNumber(), vehicle.GetVehicleColor())
	}
}
