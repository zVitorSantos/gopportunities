.PHONY: default run build rest docs clean
# Variables
APP_NAME=gopportunities

# Tasks
default: run-docs

run:
	@go run main.go
run-docs:
	@swag init
	@go run main.go

build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs