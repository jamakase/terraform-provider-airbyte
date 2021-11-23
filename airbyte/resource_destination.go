package airbyte

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jamakase/terraform-prodiver-airbyte/airbyte_sdk"
)

type resourceDestinationType struct{}

func (r resourceDestinationType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"destination_id": {
				Type:     types.StringType,
				Computed: true,
			},
			"workspace_id": {
				Type:     types.StringType,
				Required: true,
			},
			"destination_definition_id": {
				Type:     types.StringType,
				Required: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
			},
			"connection_configuration": {
				Type:     types.StringType,
				Required: true,
			},
		},
	}, nil
}

// New resource instance
func (r resourceDestinationType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceDestination{
		p: *(p.(*provider)),
	}, nil
}

type resourceDestination struct {
	p provider
}

func (r resourceDestination) Create(ctx context.Context, request tfsdk.CreateResourceRequest, response *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		response.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	var plan Destination
	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	connConfiguration, _ := unmarshalConnectorConfigJSON(plan.ConnectionConfiguration.Value)

	source, _, err := r.p.client.api.DestinationApi.CreateDestination(ctx).DestinationCreate(airbyte_sdk.DestinationCreate{
		Name:                    plan.Name.Value,
		DestinationDefinitionId: plan.DestinationDefinitionId.Value,
		WorkspaceId:             plan.WorkspaceId.Value,
		ConnectionConfiguration: connConfiguration,
	}).Execute()

	if err != nil {
		response.Diagnostics.AddError(
			"Error creating source",
			"Could not create source, unexpected error: "+err.Error(),
		)
		return
	}

	var result = Destination{
		DestinationId:           types.String{Value: source.DestinationId},
		DestinationDefinitionId: types.String{Value: source.DestinationDefinitionId},
		WorkspaceId:             types.String{Value: source.WorkspaceId},
		ConnectionConfiguration: types.String{Value: normalizeConnectorConfigJSON(source.ConnectionConfiguration)},
		Name:                    types.String{Value: source.Name},
	}

	diags = response.State.Set(ctx, result)
	response.Diagnostics.Append(diags...)

	if response.Diagnostics.HasError() {
		return
	}
}

func (r resourceDestination) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {
	// Get current state
	var state Destination
	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	sourceId := state.DestinationId.Value

	source, _, err := r.p.client.api.DestinationApi.GetDestination(ctx).DestinationIdRequestBody(airbyte_sdk.DestinationIdRequestBody{
		DestinationId: sourceId,
	}).Execute()

	if err != nil {
		response.Diagnostics.AddError(
			"Error reading source",
			"Could not read sourceID "+sourceId+": "+err.Error(),
		)
		return
	}

	state.ConnectionConfiguration = types.String{Value: normalizeConnectorConfigJSON(source.ConnectionConfiguration)}
	state.Name = types.String{Value: source.Name}

	// Set state
	diags = response.State.Set(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (r resourceDestination) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	var plan Destination
	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state Destination
	diags = request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	sourceId := state.DestinationId.Value

	connConfiguration, _ := unmarshalConnectorConfigJSON(plan.ConnectionConfiguration.Value)

	source, _, err := r.p.client.api.DestinationApi.UpdateDestination(ctx).DestinationUpdate(airbyte_sdk.DestinationUpdate{
		Name:                    plan.Name.Value,
		DestinationId:           state.DestinationId.Value,
		ConnectionConfiguration: connConfiguration,
	}).Execute()

	if err != nil {
		response.Diagnostics.AddError(
			"Error update source",
			"Could not update sourceID "+sourceId+": "+err.Error(),
		)
		return
	}

	state.ConnectionConfiguration = types.String{Value: normalizeConnectorConfigJSON(source.ConnectionConfiguration)}
	state.Name = types.String{Value: source.Name}

	// Set state
	diags = response.State.Set(ctx, state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (r resourceDestination) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	var state Destination
	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	// Get source ID from state
	sourceId := state.DestinationId.Value

	// Delete source by calling API
	_, err := r.p.client.api.DestinationApi.DeleteDestination(ctx).DestinationIdRequestBody(airbyte_sdk.DestinationIdRequestBody{DestinationId: sourceId}).Execute()
	if err != nil {
		response.Diagnostics.AddError(
			"Error deleting source",
			"Could not delete sourceId "+sourceId+": "+err.Error(),
		)
		return
	}

	// Remove resource from state
	response.State.RemoveResource(ctx)
}

func (r resourceDestination) ImportState(ctx context.Context, request tfsdk.ImportResourceStateRequest, response *tfsdk.ImportResourceStateResponse) {
	panic("implement me")
}
