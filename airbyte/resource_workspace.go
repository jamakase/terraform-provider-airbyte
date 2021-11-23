package airbyte

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jamakase/terraform-prodiver-airbyte/airbyte_sdk"
)

type resourceWorkspaceType struct{}

func (r resourceWorkspaceType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"workspace_id": {
				Type:     types.StringType,
				Computed: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
			},
			"email": {
				Type:     types.StringType,
				Optional: true,
			},
			"customer_id": {
				Type:     types.StringType,
				Computed: true,
				Optional: true,
			},
		},
	}, nil
}

// New resource instance
func (r resourceWorkspaceType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceWorkspace{
		p: *(p.(*provider)),
	}, nil
}

type resourceWorkspace struct {
	p provider
}

func (r resourceWorkspace) Create(ctx context.Context, request tfsdk.CreateResourceRequest, response *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		response.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	var plan Workspace
	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	workspace, _, err := r.p.client.api.WorkspaceApi.CreateWorkspace(ctx).WorkspaceCreate(airbyte_sdk.WorkspaceCreate{
		Name:  plan.Name.Value,
		Email: &plan.Email.Value,
	}).Execute()

	if err != nil {
		response.Diagnostics.AddError(
			"Error creating workspace",
			"Could not create workspace, unexpected error: "+err.Error(),
		)
		return
	}

	var email string

	if workspace.Email != nil {
		email = *workspace.Email
	}

	var result = Workspace{
		WorkspaceId: types.String{Value: workspace.WorkspaceId},
		Name:        types.String{Value: workspace.Name},
		Email:       types.String{Value: email, Null: workspace.Email == nil},
		CustomerId:  types.String{Value: workspace.CustomerId},
	}

	diags = response.State.Set(ctx, result)
	response.Diagnostics.Append(diags...)

	if response.Diagnostics.HasError() {
		return
	}
}

func (r resourceWorkspace) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {
	// Get current state
	var state Workspace
	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	workspaceId := state.WorkspaceId.Value

	workspace, _, err := r.p.client.api.WorkspaceApi.GetWorkspace(ctx).WorkspaceIdRequestBody(airbyte_sdk.WorkspaceIdRequestBody{
		WorkspaceId: workspaceId,
	}).Execute()

	if err != nil {
		response.Diagnostics.AddError(
			"Error reading workspace",
			"Could not read workspaceID "+workspaceId+": "+err.Error(),
		)
		return
	}

	var email string

	if workspace.Email != nil {
		email = *workspace.Email
	}

	state.Name = types.String{Value: workspace.Name}
	state.Email = types.String{Value: email, Null: workspace.Email == nil}

	// Set state
	diags = response.State.Set(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (r resourceWorkspace) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	var plan Workspace
	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state Workspace
	diags = request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	workspaceId := state.WorkspaceId.Value

	workspace, _, err := r.p.client.api.WorkspaceApi.UpdateWorkspace(ctx).WorkspaceUpdate(airbyte_sdk.WorkspaceUpdate{
		WorkspaceId:             workspaceId,
		InitialSetupComplete:    false,
		AnonymousDataCollection: false,
		News:                    false,
		SecurityUpdates:         false,
		Email:                   &plan.Email.Value,
		DisplaySetupWizard:      nil,
		Notifications:           nil,
	}).Execute()

	if err != nil {
		response.Diagnostics.AddError(
			"Error update workspace",
			"Could not update workspaceID "+workspaceId+": "+err.Error(),
		)
		return
	}

	var email string

	if workspace.Email != nil {
		email = *workspace.Email
	}

	var result = Workspace{
		WorkspaceId: types.String{Value: workspace.WorkspaceId},
		Name:        types.String{Value: workspace.Name},
		Email:       types.String{Value: email, Null: workspace.Email == nil},
		CustomerId:  types.String{Value: workspace.CustomerId},
	}

	// Set state
	diags = response.State.Set(ctx, result)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (r resourceWorkspace) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	var state Workspace
	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	// Get workspace ID from state
	workspaceId := state.WorkspaceId.Value

	// Delete workspace by calling API
	_, err := r.p.client.api.WorkspaceApi.DeleteWorkspace(ctx).WorkspaceIdRequestBody(airbyte_sdk.WorkspaceIdRequestBody{WorkspaceId: workspaceId}).Execute()
	if err != nil {
		response.Diagnostics.AddError(
			"Error deleting workspace",
			"Could not delete workspaceId "+workspaceId+": "+err.Error(),
		)
		return
	}

	// Remove resource from state
	response.State.RemoveResource(ctx)
}

func (r resourceWorkspace) ImportState(ctx context.Context, request tfsdk.ImportResourceStateRequest, response *tfsdk.ImportResourceStateResponse) {
	panic("implement me")
}
