package vehicle

import (
	"reflect"
	"testing"

	ticket "github.com/ParkingLotGolang/parking_lot/ticketservice"
)

func TestGetNewVehicle(t *testing.T) {
	// data preparation
	vCar := &Car{
		vehicleType:        VehicleTypeCar,
		registrationNumber: "KA-01-HH-9999",
		color:              "White",
		ticket:             nil,
	}
	vTruck := &Truck{
		vehicleType:        VehicleTypeTruck,
		registrationNumber: "KA-01-HH-1234",
		color:              "Blue",
		ticket:             nil,
	}

	type args struct {
		registrationNum string
		color           string
		ticket          *ticket.Ticket
		vehicleType     Type
	}
	tests := []struct {
		name string
		args args
		want Vehicle
	}{
		// TODO: Add test cases.
		{
			name: "checkIfCar",
			args: args{registrationNum: "KA-01-HH-9999", color: "White", ticket: nil, vehicleType: VehicleTypeCar},
			want: vCar,
		},
		{
			name: "checkIfTruck",
			args: args{registrationNum: "KA-01-HH-1234", color: "Blue", ticket: nil, vehicleType: VehicleTypeTruck},
			want: vTruck,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNewVehicle(tt.args.registrationNum, tt.args.color, tt.args.ticket, tt.args.vehicleType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNewVehicle() = %v, want %v", got, tt.want)
			}
		})
	}
}
