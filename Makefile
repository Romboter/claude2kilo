# Makefile for building claude2kilo

# Variables
BINARY_NAME=claude2kilo
BIN_DIR=bin
SRC=$(wildcard *.go)

# Default target
.PHONY: all
all: clean build

# Clean up previous builds
.PHONY: clean
clean:
	rm -rf $(BIN_DIR)
	mkdir -p $(BIN_DIR)

# Build for Windows
.PHONY: build-windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY_NAME).exe $(SRC)

# Build for macOS
.PHONY: build-macos
build-macos:
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY_NAME)-darwin $(SRC)

# Build for Linux
.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY_NAME)-linux $(SRC)

# Build all platforms
.PHONY: build
build: build-windows build-macos build-linux

# List generated binaries
.PHONY: list
list:
	@echo "Generated binaries:"
	ls $(BIN_DIR)
