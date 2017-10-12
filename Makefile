SHELL=bash

BUILD=build
BUILD_ARCH=$(BUILD)/$(GOOS)-$(GOARCH)
BIN_DIR?=.

export GOOS?=$(shell go env GOOS)
export GOARCH?=$(shell go env GOARCH)

build:
	@mkdir -p $(BUILD_ARCH)/$(BIN_DIR)
	go build -o $(BUILD_ARCH)/$(BIN_DIR)/dp-recipe-api main.go

debug:
	HUMAN_LOG=1 go run main.go

test:
	go test -cover $(shell go list ./... | grep -v /vendor/)

.PHONEY: test build debug
