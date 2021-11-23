package airbyte

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	airbyte_sdk "github.com/jamakase/terraform-prodiver-airbyte/airbyte_sdk"
)

type client struct {
	api *airbyte_sdk.APIClient
	//capi *smapi.Client
}

func New() tfsdk.Provider {
	return &provider{}
}

type provider struct {
	configured bool
	client     *client
}

// GetSchema
func (p *provider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"host": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
				//DefaultFunc: schema.EnvDefaultFunc("AIRBYTE_HOST", nil),
			},
		},
	}, nil
}

// Provider schema struct
type providerData struct {
	Host types.String `tfsdk:"host"`
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	// Retrieve provider data from configuration
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Warning or errors can be collected in a slice type

	abCfg := airbyte_sdk.NewConfiguration()
	abCfg.Debug = true
	abCfg.Servers[0].URL = config.Host.Value

	c := client{
		api: airbyte_sdk.NewAPIClient(abCfg),
	}

	//_, _, err := c.HealthApi.GetHealthCheck(ctx).Execute()
	//
	//if err != nil {
	//	diags = append(diags, diag.Diagnostic{
	//		Severity: diag.Error,
	//		Summary:  "Unable to create Airbyte client",
	//		Detail:   err.Error(),
	//	})
	//	return nil, diags
	//}

	p.client = &c
	p.configured = true
}

// GetResources - Defines provider resources
func (p *provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"airbyte_workspace":   resourceWorkspaceType{},
		"airbyte_source":      resourceSourceType{},
		"airbyte_destination": resourceDestinationType{},
		"airbyte_connection":  resourceConnectionType{},
	}, nil
}

// GetDataSources - Defines provider data sources
func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"airbyte_source_definition":      dataSourceSourceDefinitionType{},
		"airbyte_destination_definition": dataSourceDestinationDefinitionType{},
	}, nil
}
