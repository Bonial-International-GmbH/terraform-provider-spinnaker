package main

import (
	"github.com/armory-io/terraform-provider-spinnaker/spinnaker"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return spinnaker.Provider()
		},
	})
}
