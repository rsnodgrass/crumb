.PHONY: build install test clean run fmt lint

# Binary name
BINARY=prompt-share

# Build the binary
build:
	go build -o bin/$(BINARY) ./cmd/prompt-share

# Install to GOPATH/bin
install:
	go install ./cmd/prompt-share

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Run the app (for development)
run:
	go run ./cmd/prompt-share

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Build for multiple platforms
release:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BINARY)-darwin-amd64 ./cmd/prompt-share
	GOOS=darwin GOARCH=arm64 go build -o bin/$(BINARY)-darwin-arm64 ./cmd/prompt-share
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY)-linux-amd64 ./cmd/prompt-share

# Default target
all: fmt lint test build
