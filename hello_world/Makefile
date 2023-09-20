.PHONY: all build test update-deps clean

all: build

# Building the main application in the hello dir
build:
	cd hello && go build

# Running the tests in the hello dir
test:
	cd custom_hello && go test

# Updating the dependencies
update-deps:
	cd custom_hello && go get -u && go mod tidy
	cd hello && go get -u && go mod tidy

# Cleaning up the hello dir
clean:
	cd hello && go clean
	cd custom_hello && go clean