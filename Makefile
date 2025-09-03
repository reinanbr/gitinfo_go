.PHONY: test build clean fmt vet example help

# Default target
help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -v -cover ./...

build: ## Build the example
	cd example && go build -o gitinfo-example main.go

run: ## Run the example
	cd example && go run main.go

fmt: ## Format code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

lint: ## Run golint (requires golint to be installed)
	golint ./...

clean: ## Clean build artifacts
	cd example && rm -f gitinfo-example

check: fmt vet test ## Run all checks (format, vet, test)

deps: ## Download dependencies
	go mod download
	go mod tidy

.DEFAULT_GOAL := help
