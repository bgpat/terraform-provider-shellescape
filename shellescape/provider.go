package shellescape

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type shellescapeProvider struct {
	version string
}

func New() provider.Provider {
	return &shellescapeProvider{}
}

func (p *shellescapeProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "shellescape"
	resp.Version = p.version
}

func (p *shellescapeProvider) Schema(context.Context, provider.SchemaRequest, *provider.SchemaResponse) {
}

func (p *shellescapeProvider) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
}

func (p *shellescapeProvider) Resources(context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *shellescapeProvider) DataSources(context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		func() datasource.DataSource {
			return &QuoteDataSource{}
		},
	}
}

func (p *shellescapeProvider) Functions(context.Context) []func() function.Function {
	return []func() function.Function{
		func() function.Function {
			return &QuoteFunction{}
		},
	}
}
