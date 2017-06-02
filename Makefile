.PHONY: all start build test deps

all: deps test build

build:
	@go build

start:
	@go run main.go run

test:
	@go test ./... -v

deps:
	@go list -f '{{join .Imports "\n"}}{{"\n"}}{{join .TestImports "\n"}}' ./... | sort | uniq | grep -v apitogo | go get
