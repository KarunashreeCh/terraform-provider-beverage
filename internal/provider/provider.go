package provider

import (
	"context"

	"github.com/KarunashreeCh/terraform-provider-beverage/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &beverageProvider{}
)

// Provider model
type beverageProvider struct{}

// New initializes the provider
func New() provider.Provider {
	return &beverageProvider{}
}

// Metadata returns provider metadata
func (p *beverageProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "beverage"
}

// Schema defines the configuration schema
func (p *beverageProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Required:    true,
				Description: "Base URL for the Beverage API",
			},
		},
	}
}

// Configure initializes the provider client
func (p *beverageProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config struct {
		BaseURL types.String `tfsdk:"base_url"`
	}

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.BaseURL.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Configuration",
			"The 'base_url' must be provided.",
		)
		return
	}

	// Initialize the client
	apiClient := client.NewClient(config.BaseURL.ValueString())
	resp.DataSourceData = apiClient
	resp.ResourceData = apiClient
}

// DataSources defines the data sources implemented in the provider.
func (p *beverageProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *beverageProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewResource,
	}
}
