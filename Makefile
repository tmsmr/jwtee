all: clean test build

clean:
	rm -f jwtee

build:
	CGO_ENABLED=0 go build -ldflags "-w -s"

test:
	go test ./...

.PHONY: all clean build test
