package shellescape

import (
	"strconv"

	"github.com/alessio/shellescape"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceShellescapeQuote() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceShellescapeQuoteRead,
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

func dataSourceShellescapeQuoteRead(d *schema.ResourceData, m interface{}) error {
	quoted := shellescape.Quote(d.Get("string").(string))
	d.Set("quoted", quoted)
	d.SetId(strconv.Itoa(hashcode.String(quoted)))
	return nil
}
