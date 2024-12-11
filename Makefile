# Binary name
BINARY_NAME=slack-app

# Go related variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.PHONY: all build clean run test deps lint

all: clean build

## Build the binary
build:
	@echo "Building..."
	go build -o $(GOBIN)/$(BINARY_NAME) cmd/server/main.go

## Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

## Run the application
run:
	@echo "Running..."
	go run cmd/server/main.go

## Clean the binary
clean:
	@echo "Cleaning..."
	rm -rf $(GOBIN)
	go clean

## Run tests
test:
	@echo "Running tests..."
	go test -v ./...

## Run linter
lint:
	@echo "Running linter..."
	go vet ./...
	@if command -v golangci-lint >/dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint is not installed"; \
	fi

## Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  %-15s %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST) 