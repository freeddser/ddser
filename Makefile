#!/bin/bash

export PKGS=$(shell go list ./... | grep -v vendor/)

all:
	        build test run-server run-client
build:
	        build-server build-client

build-server:
	        @echo "Building ddser_server..."
	        @go build -o ddser_server ./cmd/server/
	        @echo "Done."


build-client:
	        @echo "Building ddser_client..."
	        @go build -o ddser_client ./cmd/client/
	        @echo "Done."


run-server:
	        @echo "Running ddser_server binary..."
	        @./ddser_server -c ./cmd/server/config.toml


run-client:
	        @echo "Running ddser_client binary..."
	        @./ddser_client -c ./cmd/client/config.toml
