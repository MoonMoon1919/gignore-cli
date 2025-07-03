# Go parameters
GOCMD=go
GOENVCMD=goenv

GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOCLEAN=$(GOCMD) clean

BINARY_NAME=gignore-cli
GOVERSION=`cat go.mod | grep 'go\s\d.' | cut -d ' ' -f2`
VERSION=`git describe --tags --always`

WINDOWS_AMD64_PATH=$(BINARY_NAME)_windowsamd64_$(VERSION).exe
LINUX_AMD64_PATH=$(BINARY_NAME)_linuxamd64_$(VERSION)
DARWIN_ARM64_PATH=$(BINARY_NAME)_darwinarm64_$(VERSION)
DARWIN_AMD64_PATH=$(BINARY_NAME)_darwinamd64_$(VERSION)

LD_FLAGS=-ldflags="-X github.com/MoonMoon1919/gignore-cli/internal/builder.VERSION=$(VERSION)" .


# Check if required tools are installed
.PHONE: check-goenv
check-goenv:
	@which $(GOENVCMD) >/dev/null 2>&1 || \
		(echo "ERROR: goenv is not installed or not in PATH" && exit 1)

.PHONY: check-tools
check-tools:
	@which $(GOCMD) >/dev/null 2>&1 || \
		(echo "ERROR: Go is not installed or not in PATH" && exit 1)

# Format all go files
.PHONY: fmt
fmt: check-tools
	@$(GOFMT) ./...

# Run go vet
.PHONY: vet
vet: check-tools
	@$(GOVET) ./...

# Download dependencies
.PHONY: deps
deps: check-tools
	@$(GOMOD) download
	@$(GOMOD) verify

# Clean build artifacts
.PHONY: clean
clean: check-tools
	@echo "Removing built binaries"
	@rm -f $(BINARY_NAME)_*
	@$(GOCLEAN)

# Run tests
.PHONY: test/unit
test/unit: check-tools
	@$(GOTEST) -v ./...

.PHONY: test/unit/cover
test/unit/cover: check-tools
	@$(GOTEST) -v -cover ./...

.PHONY: init-shell
init-shell: check-goenv
	@$(GOENVCMD) local $(GOVERSION)

# Builders
build: build/windowsamd64 build/linuxamd64 build/darwinamd64 build/darwinarm64

build/windowsamd64:
	@echo "Building binary for windows amd64"
	@GOOS=windows GOARCH=amd64 $(GOBUILD) -v -o $(WINDOWS_AMD64_PATH) $(LD_FLAGS)
	@chmod +x $(WINDOWS_AMD64_PATH)

build/linuxamd64:
	@echo "Building binary for linux amd64"
	@GOOS=linux GOARCH=amd64 $(GOBUILD) -v -o $(LINUX_AMD64_PATH) $(LD_FLAGS)
	@chmod +x $(LINUX_AMD64_PATH)

build/darwinamd64:
	@echo "Building binary for darwin amd64"
	@GOOS=darwin GOARCH=amd64 $(GOBUILD) -v -o $(DARWIN_AMD64_PATH) $(LD_FLAGS)
	@chmod +x $(DARWIN_AMD64_PATH)

build/darwinarm64:
	@echo "Building binary for darwin arm64"
	@GOOS=darwin GOARCH=arm64 $(GOBUILD) -v -o $(DARWIN_ARM64_PATH) $(LD_FLAGS)
	@chmod +x $(DARWIN_ARM64_PATH)

# Show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  clean             - Removes build artifacts"
	@echo "  deps              - Downloads and verify dependencies"
	@echo "  fmt               - Formats Go source files"
	@echo "  help              - Shows this help message"
	@echo "  test/unit         - Runs unit tests"
	@echo "  vet               - Runs go vet"
	@echo "  init-shell        - Sets goversion using goenv"
	@echo "  build             - Builds a binary for each of windows, linux, and darwin"

# Default target
.DEFAULT_GOAL := help
