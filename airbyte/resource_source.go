package airbyte

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jamakase/terraform-prodiver-airbyte/airbyte_sdk"
)

type resourceSourceType struct{}

func (r resourceSourceType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"source_id": {
				Type:     types.StringType,
				Computed: true,
			},
			"workspace_id": {
				Type:     types.StringType,
				Required: true,
			},
			"source_definition_id": {
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
func (r resourceSourceType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceSource{
		p: *(p.(*provider)),
	}, nil
}

type resourceSource struct {
	p provider
}

func (r resourceSource) Create(ctx context.Context, request tfsdk.CreateResourceRequest, response *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		response.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	var plan Source
	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	connConfiguration, _ := unmarshalConnectorConfigJSON(plan.ConnectionConfiguration.Value)

	source, _, err := r.p.client.api.SourceApi.CreateSource(ctx).SourceCreate(airbyte_sdk.SourceCreate{
		Name:                    plan.Name.Value,
		SourceDefinitionId:      plan.SourceDefinitionId.Value,
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

	var result = Source{
		SourceId:                types.String{Value: source.SourceId},
		SourceDefinitionId:      types.String{Value: source.SourceDefinitionId},
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

func (r resourceSource) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {
	// Get current state
	var state Source
	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	sourceId := state.SourceId.Value

	source, _, err := r.p.client.api.SourceApi.GetSource(ctx).SourceIdRequestBody(airbyte_sdk.SourceIdRequestBody{
		SourceId: sourceId,
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

func (r resourceSource) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	var plan Source
	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state Source
	diags = request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	sourceId := state.SourceId.Value

	connConfiguration, _ := unmarshalConnectorConfigJSON(plan.ConnectionConfiguration.Value)

	source, _, err := r.p.client.api.SourceApi.UpdateSource(ctx).SourceUpdate(airbyte_sdk.SourceUpdate{
		Name:                    plan.Name.Value,
		SourceId:                state.SourceId.Value,
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

func (r resourceSource) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	var state Source
	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	// Get source ID from state
	sourceId := state.SourceId.Value

	// Delete source by calling API
	_, err := r.p.client.api.SourceApi.DeleteSource(ctx).SourceIdRequestBody(airbyte_sdk.SourceIdRequestBody{SourceId: sourceId}).Execute()
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

func (r resourceSource) ImportState(ctx context.Context, request tfsdk.ImportResourceStateRequest, response *tfsdk.ImportResourceStateResponse) {
	panic("implement me")
}
