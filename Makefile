REPO_NAME		= github.com/ParkingLotGolang
PACKAGE_NAME	= parking_lot
SRC_DIR			= $(GOPATH)/src/${REPO_NAME}
INSTALL_PATH	= $(SRC_DIR)/${PACKAGE_NAME}/bin

#commands
GOBUILD			=	CGO_ENABLED=0 go build
GOTEST			=	go test $(shell cd \${SRC_DIR} && go list ./... | grep -v /vendor/)  -coverprofile cover.out
GOGET			=	go get
INSTALL			=	cp -rf
MKDIR			=	mkdir -p
PWD 			=	$(shell pwd)

all:: clean setup deps build test package
	echo Building in $(SRC_DIR)

build::
	cd $(SRC_DIR) ; $(GOBUILD) -o $(INSTALL_PATH)/parking_lot $(SRC_DIR)/main.go

setup::

test::
	cd $(SRC_DIR) ; $(GOTEST)

deps::
	glide cc;
	cd $(SRC_DIR) ; glide install

package::


clean::
	rm -rf $(INSTALL_PATH)/parking_lot