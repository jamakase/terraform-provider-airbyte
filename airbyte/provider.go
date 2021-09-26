package airbyte

import (
	"context"
	airbyte_sdk "github.com/jamakase/terraform-prodiver-airbyte/airbyte_sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AIRBYTE_HOST", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"airbyte_workspace": resourceWorkspace(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"airbyte_workspace": dataSourceWorkspace(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := d.Get("host").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	println(host)

	abCfg := airbyte_sdk.NewConfiguration()
	abCfg.Debug = true
	abCfg.Servers[0].URL = host

	c := airbyte_sdk.NewAPIClient(abCfg)

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

	return c, diags
}
