.PHONY: all build run clean test

all: clean build test

build:
	go build -o bin/kvstore cmd/kvstore/main.go

run:
	go run cmd/kvstore/main.go

clean:
	rm -rf bin

test:
	go test -v ./...