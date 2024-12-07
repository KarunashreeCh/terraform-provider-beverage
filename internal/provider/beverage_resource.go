package provider

import (
	"context"
	"fmt"

	"github.com/KarunashreeCh/terraform-provider-beverage/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &beverageResource{}
	_ resource.ResourceWithConfigure = &beverageResource{}
)

// type beverageResource struct{}
type beverageResource struct {
	client *client.Client
}

type beverageResourceModel struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
	ID   types.Int64  `tfsdk:"id"`
}

// NewResource initializes the beverage resource
func NewResource() resource.Resource {
	return &beverageResource{}
}

// Metadata returns resource metadata
func (r *beverageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "beverage"
}

// Schema defines the resource schema
func (r *beverageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: "ID of the beverage.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the beverage.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Type of the beverage.",
			},
		},
	}
}

func (r *beverageResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *hashicups.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Create creates a new beverage
func (r *beverageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data beverageResourceModel

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiClient := req.ProviderData.(*client.Client)

	beverage, err := apiClient.CreateBeverage(data.Name.ValueString(), data.Type.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error creating beverage", err.Error())
		return
	}

	data.ID = types.Int64Value(int64(beverage.ID))

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

// Read fetches the beverage
func (r *beverageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Similar to the Data Source Read
}

// Update updates a beverage
func (r *beverageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Implement update logic
}

// Delete deletes a beverage
func (r *beverageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Implement delete logic
}
