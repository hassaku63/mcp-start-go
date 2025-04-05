.PHONY: build clean run test lint

# Build configuration
BINARY_NAME=mcp_start
BUILD_DIR=./build
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

# Default target
all: build

# Build for current platform
build:
	@echo "Building $(BINARY_NAME) for $(GOOS)/$(GOARCH)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME)$(if $(GOOS),$(if $(filter windows,$(GOOS)),.exe,),) ./cmd/cli

# Clean
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

# Run
run: build
	@echo "Running $(BINARY_NAME)..."
	@$(BUILD_DIR)/$(BINARY_NAME)$(if $(GOOS),$(if $(filter windows,$(GOOS)),.exe,),)

# Test
test:
	@echo "Running tests..."
	go test -v -race -cover ./...

# Lint
lint:
	@echo "Running linters..."
	go vet ./...
	golangci-lint run ./...
