package shellescape

import (
	"context"

	"github.com/alessio/shellescape"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type QuoteFunction struct{}

func (f *QuoteFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "quote"
}

func (f *QuoteFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.StringParameter{
				Description: "plain string",
				Name:        "string",
			},
		},
		Return:              function.StringReturn{},
		Summary:             "Escape shell safely",
		MarkdownDescription: "Escape `string` safely in Bourne shell command line.",
	}
}

func (f *QuoteFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var s string
	resp.Error = req.Arguments.Get(ctx, &s)
	if resp.Error != nil {
		return
	}
	resp.Error = resp.Result.Set(ctx, types.StringValue(shellescape.Quote(s)))
}
