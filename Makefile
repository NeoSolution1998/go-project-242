.PHONY: build run lint test

build: 
	go build -o bin/path-size ./cmd/path-size
run:
	go run cmd/path-size/main.go

lint:
	golangci-lint run ./...

test:
	go test -v ./...
