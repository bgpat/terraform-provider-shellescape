package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/bgpat/terraform-provider-shellescape/shellescape"
)

func main() {
	err := providerserver.Serve(
		context.Background(),
		shellescape.New,
		providerserver.ServeOpts{
			Address: "registry.terraform.io/bgpat/shellescape",
		},
	)

	if err != nil {
		log.Fatal(err)
	}
}
