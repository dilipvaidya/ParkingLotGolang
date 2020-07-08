# ParkingLotGolang
This is the solution to the automated parking system implemented in Go lang.

## Problem Statement

I own a parking lot that can hold up to 'n' cars at any given point in time. Each slot is given a number starting at 1 increasing with increasing distance from the entry point in steps of one. I want to create an automated ticketing system that allows my customers to use my parking lot without human intervention.

When a car enters my parking lot, I want to have a ticket issued to the driver. The ticket issuing process includes us documenting the registration number (number plate) and the colour of the car and allocating an available parking slot to the car before actually handing over a ticket to the driver (we assume that our customers are nice enough to always park in the slots allocated to them). The customer should be allocated a parking slot which is nearest to the entry. At the exit the customer returns the ticket which then marks the slot they were using as being available.

Due to government regulation, the system should provide me with the ability to find out:
● Registration numbers of all cars of a particular colour.
● Slot number in which a car with a given registration number is parked.
● Slot numbers of all slots where a car of a particular colour is parked.

We interact with the system via a simple set of commands which produce a specific output. Please take a look at the example below, which includes all the commands you need to support - they're self explanatory. The system should allow input in two ways. Just to clarify, the same codebase should support both modes of input - we don't want two distinct submissions.
1) It should provide us with an interactive command prompt based shell where commands can be typed in
2) It should accept a filename as a parameter at the command prompt and read the commands from that file

Getting started with the setup
------

#### Systen Requirements
------
- OS - linux based system, mac
- Go lang 1.10 and above (https://golang.org/dl/)
- Makefile - system building tool
- glide - Package Management tool for go (https://glide.sh/)
- git - version control

#### Local build
Following command will install required package dependencies before going ahead with build procedure. It needs sudo access to be able to install third party package (glide for this version) in system directories. 
⋅⋅⋅If build process inetrepted with CTRL + c when asked for the password, system shows error for glide installation and proceed with building application without glide.⋅⋅⋅

###### automatic build
Automatic build process try to install glide package management tool if not already installed which may need sudo access and will prompt for password. If interespted glide installation when asked for the password, build procedure will continue without glide.
```bash
./parking_lot/bin/setup
```

###### manual build with Makefile
```bash
make build #build the code
make test #will run test cases
```

#### How to execute
1. Interactive execution 
```bash
./parking_lot/bin/parking_lot
```
2. Automated execution as per commands mentioned in the file
```bash
./parking_lot/bin/parking_lot [optional_file_input_command]
```

#### Sample output
Note: following example shows case of glide installation interupted with CTRL+C when prompted for password.

- Build Sample
```bash
$ ./parking_lot/bin/setup 
Password:
glde installation failed
cd /Users/dilip/codebase/go/src/github.com/ParkingLotGolang ; CGO_ENABLED=0 go build -o /Users/dilip/codebase/go/src/github.com/ParkingLotGolang/parking_lot/bin/parking_lot /Users/dilip/codebase/go/src/github.com/ParkingLotGolang/main.go
cd /Users/dilip/codebase/go/src/github.com/ParkingLotGolang ; go test github.com/ParkingLotGolang github.com/ParkingLotGolang/parking_lot/commandservice github.com/ParkingLotGolang/parking_lot/parkingservice github.com/ParkingLotGolang/parking_lot/parkingslot github.com/ParkingLotGolang/parking_lot/ticketservice github.com/ParkingLotGolang/parking_lot/utils github.com/ParkingLotGolang/parking_lot/vehicle  -coverprofile cover.out
?   	github.com/ParkingLotGolang	[no test files]
?   	github.com/ParkingLotGolang/parking_lot/commandservice	[no test files]
ok  	github.com/ParkingLotGolang/parking_lot/parkingservice	1.321s	coverage: 54.5% of statements
?   	github.com/ParkingLotGolang/parking_lot/parkingslot	[no test files]
?   	github.com/ParkingLotGolang/parking_lot/ticketservice	[no test files]
?   	github.com/ParkingLotGolang/parking_lot/utils	[no test files]
ok  	github.com/ParkingLotGolang/parking_lot/vehicle	1.581s	coverage: 25.0% of statements
```

- Execute Sample
```bash
$ ./parking_lot/bin/parking_lot ./parking_lot/functional_spec/fixtures/file_input.txt
Allocated slot number:  1
Allocated slot number:  2
Allocated slot number:  3
Allocated slot number:  4
Allocated slot number:  5
Allocated slot number:  6
Slot number 4 is free
Slot No.        Registration No Colour
    1           KA-01-HH-1234   White
    2           KA-01-HH-9999   White
    3           KA-01-BB-0001   Black
    5           KA-01-HH-2701   Blue
    6           KA-01-HH-3141   Black
Allocated slot number:  4
Sorry, parking lot is full
KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333
1, 2, 4
6
Not found
```

#### Future addition planned
Current solution is minimal viable solution to start ahead with. Following feattures are throught at the time of design and application can be easily extended to address those.

##### functional features
1. Backing up parking allocation onto the persistant datastore to avoid datta loss in case of system down due to any reason.
2. Add new/Delet existing parking slots to the existing parking
3. Payment charges calculation
⋅⋅1. Default charges for initial default hours of parking
⋅⋅2. changes calculation as per number of hours parking is extented
⋅⋅3. Discounts / Different charges depend on slot type (ex: handicapped people, national heros)
4. QR code of the direction on ticket to help people reach the correct parking slot
5. Advanced booking of the parking with advanced payment with limited waiting time before reallocation
6. Multi storey parking - add/delete storeys; multi-building parking - add/delette buildings

##### technical features
1. More syncronization in check In and check out operations.
2. More commands as per new functional features like calculate prize, add discount at runtime etc.
