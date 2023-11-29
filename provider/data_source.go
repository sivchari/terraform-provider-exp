package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &ExpDataSource{}

type ExpDataSource struct {
	ID    types.String `tfsdk:"id"`
	Count types.Int64  `tfsdk:"cnt"`
}

func NewDataSource() datasource.DataSource {
	return &ExpDataSource{}
}

func (m *ExpDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, res *datasource.MetadataResponse) {
	res.TypeName = fmt.Sprintf("%s_datasource", req.ProviderTypeName)
}

func (m *ExpDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, res *datasource.SchemaResponse) {
	res.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"cnt": schema.Int64Attribute{
				Optional: true,
			},
		},
	}
}

func (m *ExpDataSource) Read(ctx context.Context, req datasource.ReadRequest, res *datasource.ReadResponse) {
	var state ExpDataSource
	diags := req.Config.Get(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
	count := state.Count.ValueInt64() + 1
	state = ExpDataSource{
		ID:    state.ID,
		Count: types.Int64Value(count),
	}
	diags = res.State.Set(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}
