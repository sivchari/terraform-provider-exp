package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ provider.Provider = &ExpProvider{}

type ExpProvider struct{}

type ExpProviderConfig struct {
	Ping types.Bool `tfsdk:"ping"`
}

func New() func() provider.Provider {
	return func() provider.Provider {
		return &ExpProvider{}
	}
}

func (m *ExpProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ping": schema.BoolAttribute{
				Optional: true,
			},
		},
	}
}

func (m *ExpProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "exp"
}

func (m *ExpProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var cfg ExpProviderConfig
	diags := req.Config.Get(ctx, &cfg)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if cfg.Ping.ValueBool() {
		tflog.Info(ctx, "pong")
	}
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "configured exp-provider")
}

func (m *ExpProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewResource,
	}
}

func (m *ExpProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewDataSource,
	}
}
