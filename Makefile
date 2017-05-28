.PHONY: start build test

start:
	go run main.go run

build:
	go build

test:
	go test ./... -v
