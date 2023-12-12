PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin

GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint
LINTER_VERSION := v1.55.2

.PHONY: lint
lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run ./...
