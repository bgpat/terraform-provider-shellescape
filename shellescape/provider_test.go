package shellescape_test

import (
	"testing"

	"github.com/bgpat/terraform-provider-shellescape/shellescape"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestProvider(t *testing.T) {
	if err := shellescape.Provider().InternalValidate(); err != nil {
		t.Fatal(err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = shellescape.Provider()
}
