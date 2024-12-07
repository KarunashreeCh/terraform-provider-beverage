package provider

import (
	"context"

	"github.com/KarunashreeCh/terraform-provider-beverage/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &beverageDataSource{}
	_ datasource.DataSourceWithConfigure = &beverageDataSource{}
)

//type beverageDataSource struct{}

// NewDataSource initializes the beverage data source
func NewDataSource() datasource.DataSource {
	return &beverageDataSource{}
}

type beverageDataSource struct {
	client *client.Client
}
type beverageDataSourceModel struct {
	Beverage []beverageResourceModel `tfsdk:"beverages"`
}



// Metadata returns data source metadata
func (d *beverageDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "beverage"
}

// Schema defines the data source schema
func (d *beverageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required:    true,
				Description: "ID of the beverage to fetch.",
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the beverage.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "Type of the beverage.",
			},
		},
	}
}

// Read fetches the beverage data
func (d *beverageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data struct {
		ID   types.Int64  `tfsdk:"id"`
		Name types.String `tfsdk:"name"`
		Type types.String `tfsdk:"type"`
	}

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiClient := req.ProviderData.(*client.Client)

	beverage, err := apiClient.GetBeverage(int(data.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError("Error fetching beverage", err.Error())
		return
	}

	data.Name = types.StringValue(beverage.Name)
	data.Type = types.StringValue(beverage.Type)

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}
