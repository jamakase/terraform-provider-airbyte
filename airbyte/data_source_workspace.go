package airbyte

import (
	"context"
	"github.com/jamakase/terraform-prodiver-airbyte/airbyte_sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWorkspace() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorkspaceRead,
		Schema: map[string]*schema.Schema{
			"workspace_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceWorkspaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*airbyte_sdk.APIClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	workspaceId := d.Get("workspace_id").(string)

	workspace, _, err := c.WorkspaceApi.GetWorkspace(ctx).WorkspaceIdRequestBody(airbyte_sdk.WorkspaceIdRequestBody{
		WorkspaceId: workspaceId,
	}).Execute()

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(workspace.WorkspaceId)
	d.Set("name", workspace.Name)
	d.Set("workspace_id", workspace.WorkspaceId)

	return diags
}
