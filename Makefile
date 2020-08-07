export GO111MODULE=on
VERSION=$(shell git describe --tags --candidates=1 --dirty)
BUILD_FLAGS=-ldflags="-X main.version=$(VERSION)"
SRC=$(shell find . -name '*.go')

.PHONY: all clean release install

all: linux

clean:
	rm -f drone-teams linux

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...

linux: $(SRC)
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o drone-teams-linux .

install:
	rm -f drone-teams
	go build $(BUILD_FLAGS) .
	mv drone-teams ~/bin/
