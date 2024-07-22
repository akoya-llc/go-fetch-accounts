#!/bin/bash

# Step 1: Download all dependencies listed in go.mod
#go mod download
#
#go get -u -v -f all
go mod tidy
go mod vendor

# Step 2: Build the main.go file
go build -o main .

# Step 3: Run the main executable
./main
