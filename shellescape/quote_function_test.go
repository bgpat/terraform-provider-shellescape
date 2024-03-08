package shellescape_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestQuoteFunction(t *testing.T) {
	for k, v := range chars {
		t.Run(k, func(t *testing.T) {
			resource.ParallelTest(t, resource.TestCase{
				IsUnitTest:               true,
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: fmt.Sprintf(`output "test" { value = provider::shellescape::quote("%s") }`, k),
						Check:  resource.ComposeTestCheckFunc(resource.TestCheckOutput("test", v)),
					},
				},
			})
		})
	}
}
