package airbyte

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	airbyte_sdk "github.com/jamakase/terraform-prodiver-airbyte/airbyte_sdk"
)

func resourceWorkspace() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkspaceCreate,
		ReadContext:   resourceWorkspaceRead,
		UpdateContext: resourceWorkspaceUpdate,
		DeleteContext: resourceWorkspaceDelete,
		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		//Importer: &schema.ResourceImporter{
		//	State: schema.ImportStatePassthroughContext,
		//},
	}
}

func resourceWorkspaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*airbyte_sdk.APIClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	workspace, _, err := c.WorkspaceApi.CreateWorkspace(ctx).WorkspaceCreate(airbyte_sdk.WorkspaceCreate{
		Name: d.Get("name").(string),
	}).Execute()

	println(workspace.WorkspaceId)
	if err != nil {
		diag.FromErr(err)
	}

	d.SetId(workspace.WorkspaceId)

	return diags
}

func resourceWorkspaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*airbyte_sdk.APIClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	workspaceId := d.Id()

	workspace, _, err := c.WorkspaceApi.GetWorkspace(ctx).WorkspaceIdRequestBody(airbyte_sdk.WorkspaceIdRequestBody{
		workspaceId,
	}).Execute()

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(workspace.WorkspaceId)
	d.Set("name", workspace.Name)
	d.Set("workspace_id", workspace.WorkspaceId)

	return diags
}

func resourceWorkspaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//c := m.(*airbyte_sdk.APIClient)

	return resourceWorkspaceRead(ctx, d, m)
}

func resourceWorkspaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*airbyte_sdk.APIClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req := c.WorkspaceApi.DeleteWorkspace(ctx)

	workspaceId := d.Get("id").(string)
	req.WorkspaceIdRequestBody(airbyte_sdk.WorkspaceIdRequestBody{
		WorkspaceId: workspaceId,
	})

	_, err := req.Execute()

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
