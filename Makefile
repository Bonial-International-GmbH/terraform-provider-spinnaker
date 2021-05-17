VERSION         = 0.0.3
LOCAL_PROVIDERS ="$$HOME/.terraform.d/plugins_local"
BINARY_PATH     = "registry.terraform.io/Bonial-International-GmbH/spinnaker/${VERSION}/$$(go env GOOS)_$$(go env GOARCH)/terraform-provider-spinnaker_${VERSION}"

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
build_local:
	@echo "Please configure your .terraformrc file to contain a filesystem_mirror block pointed at '${LOCAL_PROVIDERS}' for 'registry.terraform.io/Bonial-International-GmbH/spinnaker'"
	@echo "You MUST use a direct exclusion in this block in order to pick up the built binary. Otherwise it will query the registry and find a distribution version"
	@echo "You MUST comment out the 'version' constraint in the required_providers block in any Terraform installation you test this in."
	@echo "You MUST delete existing cached plugins from any .terraform directories in Terraform installations you want to test against so that it will perform a lookup on the local mirror"
	go build -o "${LOCAL_PROVIDERS}/${BINARY_PATH}"
