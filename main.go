package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/bgpat/terraform-provider-shellescape/shellescape"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: shellescape.Provider,
	})
}
