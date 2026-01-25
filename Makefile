# Makefile for claude-test

# Variables
APP_NAME := claude-test
GO := go

.PHONY: all build run test tidy up down clean help

# Default target
all: run

# Build the binary
build:
	$(GO) build -o bin/$(APP_NAME) ./...

# Run the application
run:
	$(GO) run main.go

# Run tests
test:
	$(GO) test ./...

# Tidy dependencies
 tidy:
	$(GO) mod tidy

# Start MariaDB using docker compose
up:
	docker compose -f docker-compose.yaml up -d

# Stop MariaDB
 down:
	docker compose -f docker-compose.yaml down

# Clean up
clean:
	rm -rf bin

help:
	@echo "Available targets:"
	@echo "  build   Build the binary"
	@echo "  run     Run the application"
	@echo "  test    Run tests"
	@echo "  tidy    Tidy go.mod"
	@echo "  up      Start MariaDB"
	@echo "  down    Stop MariaDB"
	@echo "  clean   Remove build artifacts"
	@echo "  help    Show this message"
