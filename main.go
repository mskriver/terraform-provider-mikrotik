package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/mskriver/terraform-provider-mikrotik/mikrotik"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mikrotik.Provider,
	})
}
