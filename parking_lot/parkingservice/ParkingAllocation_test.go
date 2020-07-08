package parkingservice

import (
	"reflect"
	"testing"

	slot "github.com/ParkingLotGolang/parking_lot/parkingslot"
	vehicle "github.com/ParkingLotGolang/parking_lot/vehicle"
)

func TestCreateParking(t *testing.T) {
	type args struct {
		totalNumberOfSlots int
	}

	// data preparation
	p := CreateParking(10)

	tests := []struct {
		name string
		args args
		want *ParkingService
	}{
		// TODO: Add test cases.
		{name: "singleton class", args: args{totalNumberOfSlots: 5}, want: p},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateParking(tt.args.totalNumberOfSlots); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateParking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingService_GetNumberOfSlots(t *testing.T) {
	// data preparation
	totalParkingSlots := 10
	p := CreateParking(totalParkingSlots)

	type fields struct {
		numberOfSlots           int
		slotAllocation          TypeSlotAllocationMap
		slotVehicleAllocation   TypeSlotVehicleAllocationMap
		vehicleToSlotAllocation TypeCheckedInVehiclesMap
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "singleton class",
			fields: fields{
				numberOfSlots:           totalParkingSlots + 1,
				slotAllocation:          p.GetSlotAllocationMap(),
				slotVehicleAllocation:   p.GetSlotVehicleAllocation(),
				vehicleToSlotAllocation: p.GetVehicleToSlotAllocation(),
			},
			want: totalParkingSlots,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ParkingService{
				numberOfSlots:           tt.fields.numberOfSlots,
				slotAllocation:          tt.fields.slotAllocation,
				slotVehicleAllocation:   tt.fields.slotVehicleAllocation,
				vehicleToSlotAllocation: tt.fields.vehicleToSlotAllocation,
			}
			if got := p.GetNumberOfSlots(); got != tt.want {
				t.Errorf("ParkingService.GetNumberOfSlots() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingService_CheckInNewVehicle(t *testing.T) {
	// data preparation
	totalParkingSlots := 1
	p := CreateParking(totalParkingSlots)

	type fields struct {
		numberOfSlots           int
		slotAllocation          TypeSlotAllocationMap
		slotVehicleAllocation   TypeSlotVehicleAllocationMap
		vehicleToSlotAllocation TypeCheckedInVehiclesMap
	}
	type args struct {
		licenseNumber string
		color         string
		vehicleType   vehicle.Type
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Checkin First vehicle",
			fields: fields{
				numberOfSlots:           totalParkingSlots + 1,
				slotAllocation:          p.GetSlotAllocationMap(),
				slotVehicleAllocation:   p.GetSlotVehicleAllocation(),
				vehicleToSlotAllocation: p.GetVehicleToSlotAllocation(),
			},
			args:    args{licenseNumber: "KA-01-BB-0001", color: "Black", vehicleType: vehicle.VehicleTypeCar},
			wantErr: true,
		},
		{
			name: "checkin second vehicle no space",
			fields: fields{
				numberOfSlots:           0,
				slotAllocation:          p.GetSlotAllocationMap(),
				slotVehicleAllocation:   p.GetSlotVehicleAllocation(),
				vehicleToSlotAllocation: p.GetVehicleToSlotAllocation(),
			},
			args:    args{licenseNumber: "KA-01-HH-1234", color: "White", vehicleType: vehicle.VehicleTypeCar},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ParkingService{
				numberOfSlots:           tt.fields.numberOfSlots,
				slotAllocation:          tt.fields.slotAllocation,
				slotVehicleAllocation:   tt.fields.slotVehicleAllocation,
				vehicleToSlotAllocation: tt.fields.vehicleToSlotAllocation,
			}
			if err := p.CheckInNewVehicle(tt.args.licenseNumber, tt.args.color, tt.args.vehicleType); (err == nil) != tt.wantErr {
				t.Errorf("ParkingService.CheckInNewVehicle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParkingService_CheckOutVehicle(t *testing.T) {
	// data preparation
	totalParkingSlots := 1
	p := CreateParking(totalParkingSlots)

	type fields struct {
		numberOfSlots           int
		slotAllocation          TypeSlotAllocationMap
		slotVehicleAllocation   TypeSlotVehicleAllocationMap
		vehicleToSlotAllocation TypeCheckedInVehiclesMap
	}
	type args struct {
		sltNum int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Checkin First vehicle",
			fields: fields{
				numberOfSlots:           totalParkingSlots + 1,
				slotAllocation:          p.GetSlotAllocationMap(),
				slotVehicleAllocation:   p.GetSlotVehicleAllocation(),
				vehicleToSlotAllocation: p.GetVehicleToSlotAllocation(),
			},
			args:    args{sltNum: 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ParkingService{
				numberOfSlots:           tt.fields.numberOfSlots,
				slotAllocation:          tt.fields.slotAllocation,
				slotVehicleAllocation:   tt.fields.slotVehicleAllocation,
				vehicleToSlotAllocation: tt.fields.vehicleToSlotAllocation,
			}
			if err := p.CheckOutVehicle(tt.args.sltNum); (err == nil) != tt.wantErr {
				t.Errorf("ParkingService.CheckOutVehicle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParkingService_GetNextFreeParkingSlot(t *testing.T) {
	// data preparation
	totalParkingSlots := 5
	p := CreateParking(totalParkingSlots)
	p.GetSlotAllocationMap()[1].SetSlotStatus(slot.Occupied)
	p.GetSlotAllocationMap()[2].SetSlotStatus(slot.Occupied)
	p.GetSlotAllocationMap()[3].SetSlotStatus(slot.Occupied)
	p.GetSlotAllocationMap()[4].SetSlotStatus(slot.Free)
	p.GetSlotAllocationMap()[5].SetSlotStatus(slot.Occupied)

	type fields struct {
		numberOfSlots           int
		slotAllocation          TypeSlotAllocationMap
		slotVehicleAllocation   TypeSlotVehicleAllocationMap
		vehicleToSlotAllocation TypeCheckedInVehiclesMap
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "allocation in the middle",
			fields: fields{
				numberOfSlots:           totalParkingSlots + 1,
				slotAllocation:          p.GetSlotAllocationMap(),
				slotVehicleAllocation:   p.GetSlotVehicleAllocation(),
				vehicleToSlotAllocation: p.GetVehicleToSlotAllocation(),
			},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ParkingService{
				numberOfSlots:           tt.fields.numberOfSlots,
				slotAllocation:          tt.fields.slotAllocation,
				slotVehicleAllocation:   tt.fields.slotVehicleAllocation,
				vehicleToSlotAllocation: tt.fields.vehicleToSlotAllocation,
			}
			got, err := p.GetNextFreeParkingSlot()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkingService.GetNextFreeParkingSlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParkingService.GetNextFreeParkingSlot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingService_FreeParkingSlot(t *testing.T) {
	// data preparation
	totalParkingSlots := 5
	p := CreateParking(totalParkingSlots)
	p.GetSlotAllocationMap()[1].SetSlotStatus(slot.Occupied)
	p.GetSlotAllocationMap()[2].SetSlotStatus(slot.NotInUsed)
	p.GetSlotAllocationMap()[3].SetSlotStatus(slot.Occupied)
	p.GetSlotAllocationMap()[4].SetSlotStatus(slot.Free)
	p.GetSlotAllocationMap()[5].SetSlotStatus(slot.Occupied)

	type fields struct {
		numberOfSlots           int
		slotAllocation          TypeSlotAllocationMap
		slotVehicleAllocation   TypeSlotVehicleAllocationMap
		vehicleToSlotAllocation TypeCheckedInVehiclesMap
	}
	type args struct {
		slotNumber int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "free already free slot",
			fields: fields{
				numberOfSlots:           totalParkingSlots + 1,
				slotAllocation:          p.GetSlotAllocationMap(),
				slotVehicleAllocation:   p.GetSlotVehicleAllocation(),
				vehicleToSlotAllocation: p.GetVehicleToSlotAllocation(),
			},
			args:    args{slotNumber: 4},
			wantErr: false,
		},
		{
			name: "free allocated slot",
			fields: fields{
				numberOfSlots:           totalParkingSlots + 1,
				slotAllocation:          p.GetSlotAllocationMap(),
				slotVehicleAllocation:   p.GetSlotVehicleAllocation(),
				vehicleToSlotAllocation: p.GetVehicleToSlotAllocation(),
			},
			args:    args{slotNumber: 5},
			wantErr: true,
		},
		{
			name: "free slot that is not in use",
			fields: fields{
				numberOfSlots:           totalParkingSlots + 1,
				slotAllocation:          p.GetSlotAllocationMap(),
				slotVehicleAllocation:   p.GetSlotVehicleAllocation(),
				vehicleToSlotAllocation: p.GetVehicleToSlotAllocation(),
			},
			args:    args{slotNumber: 2},
			wantErr: false,
		},
		{
			name: "free slot that does not exist",
			fields: fields{
				numberOfSlots:           totalParkingSlots + 1,
				slotAllocation:          p.GetSlotAllocationMap(),
				slotVehicleAllocation:   p.GetSlotVehicleAllocation(),
				vehicleToSlotAllocation: p.GetVehicleToSlotAllocation(),
			},
			args:    args{slotNumber: 20},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ParkingService{
				numberOfSlots:           tt.fields.numberOfSlots,
				slotAllocation:          tt.fields.slotAllocation,
				slotVehicleAllocation:   tt.fields.slotVehicleAllocation,
				vehicleToSlotAllocation: tt.fields.vehicleToSlotAllocation,
			}
			if err := p.FreeParkingSlot(tt.args.slotNumber); (err == nil) != tt.wantErr {
				t.Errorf("ParkingService.FreeParkingSlot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
