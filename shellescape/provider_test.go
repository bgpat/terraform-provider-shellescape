package shellescape_test

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	"github.com/bgpat/terraform-provider-shellescape/shellescape"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"shellescape": providerserver.NewProtocol6WithError(shellescape.New()),
}
