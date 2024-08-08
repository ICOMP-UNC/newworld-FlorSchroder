#!/bin/bash

echo "Initializing the swag..."

swag init -g cmd/main.go -o ./docs --parseDependency

echo "Starting the server..."
go run cmd/main.go
