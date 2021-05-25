.DEFAULT_GOAL := help

BINARY          := terraform-provider-spinnaker
TEST_FLAGS      ?= -race
PKGS            ?= $(shell go list ./... | grep -v /vendor/)
VERSION         ?= 99.99.99
LOCAL_PROVIDERS ="$$HOME/.terraform.d/plugins_local"
BINARY_PATH     = "registry.terraform.io/Bonial-International-GmbH/spinnaker/${VERSION}/$$(go env GOOS)_$$(go env GOARCH)/${BINARY}_${VERSION}"

.PHONY: help
help:
	@grep -E '^[a-zA-Z-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "[32m%-12s[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## build terraform-provider-spinnaker
	go build \
		-ldflags "-s -w" \
		-o $(BINARY) \
		main.go

# Builds the provider and adds it to an independently configured filesystem_mirror folder.
#
# For this to work your local ~/.terraformrc has to include the following:
#
# provider_installation {
#   filesystem_mirror {
#     path    = "<your-home-directory>/.terraform.d/plugins_local/"
#     include = ["registry.terraform.io/Bonial-International-GmbH/spinnaker"]
#   }
#
#   direct {
#     exclude = ["registry.terraform.io/Bonial-International-GmbH/spinnaker"]
#   }
# }
#
.PHONY: build_local
build_local: ## build terraform-provider-spinnaker and install it in a local plugins dir
	@echo "Please configure your .terraformrc file to contain a filesystem_mirror block pointed at '${LOCAL_PROVIDERS}' for 'registry.terraform.io/Bonial-International-GmbH/spinnaker'"
	@echo "You MUST use a direct exclusion in this block in order to pick up the built binary. Otherwise it will query the registry and find a distribution version"
	@echo "You MUST comment out the 'version' constraint in the required_providers block in any Terraform installation you test this in."
	@echo "You MUST delete existing cached plugins from any .terraform directories in Terraform installations you want to test against so that it will perform a lookup on the local mirror"
	go build -o "${LOCAL_PROVIDERS}/${BINARY_PATH}"

.PHONY: test
test: ## run tests
	go test $(TEST_FLAGS) $(PKGS)

.PHONY: vet
vet: ## run go vet
	go vet $(PKGS)

.PHONY: coverage
coverage: ## generate code coverage
	go test $(TEST_FLAGS) -covermode=atomic -coverprofile=coverage.txt $(PKGS)
	go tool cover -func=coverage.txt

.PHONY: lint
lint: ## run golangci-lint
	golangci-lint run
