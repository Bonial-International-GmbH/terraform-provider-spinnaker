package main

import (
	"github.com/Bonial-International-GmbH/terraform-provider-spinnaker/spinnaker"
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
