package shellescape

import (
	"context"
	"hash/crc32"
	"strconv"

	"github.com/alessio/shellescape"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceShellescapeQuote() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceShellescapeQuoteRead,
		Schema: map[string]*schema.Schema{
			"string": {
				Type:     schema.TypeString,
				Required: true,
			},
			"quoted": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceShellescapeQuoteRead(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	quoted := shellescape.Quote(d.Get("string").(string))
	if err := d.Set("quoted", quoted); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.Itoa(hashcode(quoted)))
	return nil
}

// https://developer.hashicorp.com/terraform/plugin/sdkv2/guides/v2-upgrade-guide#removal-of-helper-hashcode-package
func hashcode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}
