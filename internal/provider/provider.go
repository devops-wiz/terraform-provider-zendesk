// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	"github.com/JacobPotter/go-zendesk/credentialtypes"
	"github.com/JacobPotter/go-zendesk/zendesk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.Provider = &ZendeskProvider{}

// ZendeskProvider defines the provider implementation.
type ZendeskProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// ZendeskProviderModel describes the provider data model.
type ZendeskProviderModel struct {
	Subdomain types.String `tfsdk:"subdomain"`
	Username  types.String `tfsdk:"username"`
	APIToken  types.String `tfsdk:"api_token"`
}

func (p *ZendeskProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "zendesk"
	resp.Version = p.version
}

func (p *ZendeskProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with Zendesk.",
		Attributes: map[string]schema.Attribute{
			"subdomain": schema.StringAttribute{
				Description: "URI for Zendesk API. May also be provided via ZENDESK_SUBDOMAIN environment variable.",
				Optional:    true,
			},
			"username": schema.StringAttribute{
				Description: "Username for Zendesk API. May also be provided via ZENDESK_USERNAME environment variable.",
				Optional:    true,
			},
			"api_token": schema.StringAttribute{
				Description: "APIToken for Zendesk API. May also be provided via ZENDESK_PASSWORD environment variable.",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

func (p *ZendeskProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config ZendeskProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Subdomain.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("subdomain"),
			"Unknown Zendesk API Host",
			"The provider cannot create the Zendesk API client as there is an unknown configuration value for the Zendesk API subdomain. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the ZENDESK_SUBDOMAIN environment variable.",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown Zendesk API Username",
			"The provider cannot create the Zendesk API client as there is an unknown configuration value for the Zendesk API username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the ZENDESK_USERNAME environment variable.",
		)
	}

	if config.APIToken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_token"),
			"Unknown Zendesk API APIToken",
			"The provider cannot create the Zendesk API client as there is an unknown configuration value for the Zendesk API api_token. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the ZENDESK_PASSWORD environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	subdomain := os.Getenv("ZENDESK_SUBDOMAIN")
	username := os.Getenv("ZENDESK_USERNAME")
	apiToken := os.Getenv("ZENDESK_API_TOKEN")

	if !config.Subdomain.IsNull() {
		subdomain = config.Subdomain.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.APIToken.IsNull() {
		apiToken = config.APIToken.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	emptyValueMessage := "If either is already set, ensure the value is not empty."

	if subdomain == "" {

		resp.Diagnostics.AddAttributeError(
			path.Root("subdomain"),
			"Missing Zendesk API Host",
			"The provider cannot create the Zendesk API client as there is a missing or empty value for the Zendesk API subdomain. "+
				"Set the subdomain value in the configuration or use the ZENDESK_SUBDOMAIN environment variable. "+
				emptyValueMessage,
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing Zendesk API Username",
			"The provider cannot create the Zendesk API client as there is a missing or empty value for the Zendesk API username. "+
				"Set the username value in the configuration or use the ZENDESK_USERNAME environment variable. "+
				emptyValueMessage,
		)
	}

	if apiToken == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_token"),
			"Missing Zendesk API APIToken",
			"The provider cannot create the Zendesk API client as there is a missing or empty value for the Zendesk API api_token. "+
				"Set the api_token value in the configuration or use the ZENDESK_API_TOKEN environment variable. "+
				emptyValueMessage,
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "zendesk_subdomain", subdomain)
	ctx = tflog.SetField(ctx, "zendesk_username", username)
	ctx = tflog.SetField(ctx, "zendesk_api_token", apiToken)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "zendesk_api_token")

	tflog.Info(ctx, "Creating Zendesk client")

	// Create a new Zendesk client using the configuration values
	// client, err := client.NewClient(&subdomain, &username, &apiToken)
	client, err := zendesk.NewClient(nil)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Zendesk API Client",
			"An unexpected error occurred when creating the Zendesk API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Zendesk Client Error: "+err.Error(),
		)
		return
	}

	err = client.SetSubdomain(subdomain)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Zendesk API Client",
			"An unexpected error occurred when creating the Zendesk API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Zendesk Client Error: "+err.Error(),
		)
		return
	}

	client.SetCredential(credentialtypes.NewAPITokenCredential(username, apiToken))

	// Make the Zendesk client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured Zendesk client", map[string]any{"success": true})

}

func (p *ZendeskProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewTicketFieldResource,
		NewMacroResource,
		NewTriggerCategoryResource,
		NewTriggerResource,
		NewAutomationResource,
		NewViewResource,
		NewWebhookResource,
		NewSLAResource,
		NewTicketFormResource,
		NewGroupResource,
		NewBrandResource,
		NewUserFieldResource,
		NewOrganizationFieldResource,
		NewScheduleResource,
		NewDynamicContentResource,
	}
}

func (p *ZendeskProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewSearchDatasource,
		NewLocaleDatasource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ZendeskProvider{
			version: version,
		}
	}
}
