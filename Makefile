GO = go

APP_NAME := mars-rover-simulator

.PHONY: build run test clean fmt coverage

build:
	go build -o $(APP_NAME) .

run: build
	./$(APP_NAME)

test:
	go test -v ./... -coverprofile=coverage.out 

coverage:
	go tool cover -html=coverage.out -o coverage.html

fmt:
	go fmt ./...

clean:
	rm -f $(APP_NAME)