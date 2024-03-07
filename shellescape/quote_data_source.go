package shellescape

import (
	"context"

	"github.com/alessio/shellescape"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type QuoteDataSource struct{}

type QuoteDataSourceModel struct {
	String types.String `tfsdk:"string"`
	Quoted types.String `tfsdk:"quoted"`
}

func (d *QuoteDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_quote"
}

func (d *QuoteDataSource) Schema(_ context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"string": schema.StringAttribute{
				Required: true,
			},
			"quoted": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *QuoteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data QuoteDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	data.Quoted = types.StringValue(shellescape.Quote(data.String.ValueString()))
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
