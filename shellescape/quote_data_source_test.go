package shellescape_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestQuoteDataSource(t *testing.T) {
	for k, v := range chars {
		t.Run(k, func(t *testing.T) {
			resource.ParallelTest(t, resource.TestCase{
				IsUnitTest:               true,
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: fmt.Sprintf(`data "shellescape_quote" "test" { string = "%s" }`, k),
						Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("data.shellescape_quote.test", "quoted", v)),
					},
				},
			})
		})
	}
}
