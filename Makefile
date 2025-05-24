APP := Mongo-Server
BIN := bin
MAIN := main.go

.PHONY: all build clean

## all: Default target (same as build)
all: build

## build: Builds the go binary and stores in bin/
build:
	@echo "Building $(APP)..."
	@mkdir $(BIN)
	go build -o $(BIN)/$(APP) $(MAIN)
	@echo "Build Completed!"

## clean: Cleans previous builds
clean:
	@echo "Cleaning previous build of $(APP)..."
	@rm -rf $(BIN)
	@echo "Clean Completed!"

## help: Show this message
help:
	@echo "Available Commands:"
	@grep -E '^##' $(MAKEFILE_LIST)| awk -F ': ' '{printf "%-15s %s\n",$$1,$$2}' | sed -E 's/^## //' | sort