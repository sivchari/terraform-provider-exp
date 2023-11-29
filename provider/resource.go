package provider

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = &ExpResource{}

type ExpResource struct {
	ID    types.String `tfsdk:"id"`
	Count types.Int64  `tfsdk:"cnt"`
}

func NewResource() resource.Resource {
	return &ExpResource{}
}

func (m *ExpResource) Metadata(_ context.Context, req resource.MetadataRequest, res *resource.MetadataResponse) {
	res.TypeName = fmt.Sprintf("%s_resource", req.ProviderTypeName)
}

func (m *ExpResource) Schema(_ context.Context, _ resource.SchemaRequest, res *resource.SchemaResponse) {
	res.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cnt": schema.Int64Attribute{
				Computed: true,
			},
		},
	}
}

func (m *ExpResource) Create(ctx context.Context, req resource.CreateRequest, res *resource.CreateResponse) {
	var state ExpResource
	diags := req.Plan.Get(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
	count := state.Count.ValueInt64() + 1
	id := uuid.New()
	state = ExpResource{
		ID:    types.StringValue(id.String()),
		Count: types.Int64Value(count),
	}
	diags = res.State.Set(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}

func (m *ExpResource) Read(ctx context.Context, req resource.ReadRequest, res *resource.ReadResponse) {
	var state ExpResource
	diags := req.State.Get(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
	diags = res.State.Set(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}

func (m *ExpResource) Update(ctx context.Context, req resource.UpdateRequest, res *resource.UpdateResponse) {
	var state ExpResource
	diags := req.Plan.Get(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
	count := state.Count.ValueInt64() + 1
	state = ExpResource{
		ID:    state.ID,
		Count: types.Int64Value(count),
	}
	diags = res.State.Set(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}

func (m *ExpResource) Delete(ctx context.Context, req resource.DeleteRequest, res *resource.DeleteResponse) {
	var state ExpResource
	diags := req.State.Get(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "start to delete resource")
	res.State.RemoveResource(ctx)
}
