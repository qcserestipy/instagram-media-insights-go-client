SHELL := /usr/bin/env bash

.DEFAULT_GOAL := help

DOCKERCMD=$(shell which docker)
SWAGGER_VERSION=v0.30.3
SWAGGER := $(DOCKERCMD) run --rm -t -u "$(shell id -u):$(shell id -g)" -v $(shell pwd):/src -w /src quay.io/goswagger/swagger:$(SWAGGER_VERSION)

ifeq ($(VERSION),)
VERSION := v2.13.1
endif

API_SPEC=api/v24.0/swagger.yaml
API_SPEC_ACCOUNT=api/v24.0/swagger-account.yaml
CLIENT_DIR=pkg/sdk/v24.0
CLIENT_DIR_ACCOUNT=pkg/sdk-account/v24.0

## --------------------------------------
## Help
## --------------------------------------

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## Client
## --------------------------------------

.PHONY: gen-api-client
gen-api-client: ## Generate goswagger client for media insights
	@$(SWAGGER) generate client -f ${API_SPEC} --target=$(CLIENT_DIR) --template=stratoscale --additional-initialism=CVE --additional-initialism=GC --additional-initialism=OIDC

.PHONY: gen-api-client-account
gen-api-client-account: ## Generate goswagger client for account insights
	@$(SWAGGER) generate client -f ${API_SPEC_ACCOUNT} --target=$(CLIENT_DIR_ACCOUNT) --template=stratoscale --additional-initialism=CVE --additional-initialism=GC --additional-initialism=OIDC

.PHONY: gen-all-clients
gen-all-clients: gen-api-client gen-api-client-account ## Generate all API clients

.PHONY: build
build: ## Build the main application
	go build -o bin/instagram-media-insights-go-client cmd/main/main.go

.PHONY: cleanup
cleanup: ## Clean up generated client code
	rm -rf pkg/sdk/v24.0/models pkg/sdk/v24.0/client pkg/sdk-account/v24.0/models pkg/sdk-account/v24.0/client
.PHONY: test
test: ## run the test
	go test ./...

.PHONY: update-submodule
update-submodule:
	@echo "Updating api submodule..."
	@cd api && git stash && git pull --rebase && cd ..
	@echo "Submodule updated successfully"

get-go-version:
	@echo "$(VERSION)" | sed -E "s/v([0-9]+)\.([0-9]+)\.([0-9]+)/v0.\1\2.\3/g" -

all: cleanup gen-all-clients