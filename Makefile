BUILD_TIME := $(shell date -u '+%Y-%m-%d %H:%M:%S')
VERSION := 0.1.0
build:
	@go build -ldflags="-X main.Version=$(VERSION) -X 'main.BuildTime=$(BUILD_TIME)'" -o bin/cool-cli

run: build
	@./bin/cool-cli
