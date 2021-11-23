package airbyte

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jamakase/terraform-prodiver-airbyte/airbyte_sdk"
)

type resourceConnectionType struct{}

func (r resourceConnectionType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"connection_id": {
				Type:     types.StringType,
				Computed: true,
			},
			"source_id": {
				Type:     types.StringType,
				Required: true,
			},
			"destination_id": {
				Type:     types.StringType,
				Required: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
			},
			"schedule": {
				Optional: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"units": {
						Type:     types.Int64Type,
						Required: true,
					},
					"time_unit": {
						Type:     types.StringType,
						Required: true,
					},
				}),
			},
			"status": {
				Type:     types.StringType,
				Optional: true,
			},
		},
	}, nil
}

// New resource instance
func (r resourceConnectionType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceConnection{
		p: *(p.(*provider)),
	}, nil
}

type resourceConnection struct {
	p provider
}

func (r resourceConnection) Create(ctx context.Context, request tfsdk.CreateResourceRequest, response *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		response.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	var plan Connection
	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	connection, _, err := r.p.client.api.ConnectionApi.CreateConnection(ctx).ConnectionCreate(airbyte_sdk.ConnectionCreate{
		Name: &plan.Name.Value,
		//NamespaceDefinition:  nil,
		//NamespaceFormat:      nil,
		//Prefix:               nil,
		SourceId:      plan.SourceId.Value,
		DestinationId: plan.DestinationId.Value,
		//OperationIds:         nil,
		//SyncCatalog:          nil,
		//Schedule:             nil,
		Status: plan.Status,
		Schedule: &airbyte_sdk.ConnectionSchedule{
			Units:    plan.Schedule.Units.Value,
			TimeUnit: plan.Schedule.TimeUnit.Value,
		},
		//ResourceRequirements: nil,
	}).Execute()

	if err != nil {
		response.Diagnostics.AddError(
			"Error creating source",
			"Could not create source, unexpected error: "+err.Error(),
		)
		return
	}

	var schedule *Schedule

	if connection.Schedule != nil {
		schedule = &Schedule{
			Units: types.Int64{
				Value: connection.Schedule.Units,
			},
			TimeUnit: types.String{
				Value: connection.Schedule.TimeUnit,
			},
		}
	}

	var result = Connection{
		ConnectionId:  types.String{Value: connection.ConnectionId},
		SourceId:      types.String{Value: connection.SourceId},
		DestinationId: types.String{Value: connection.DestinationId},
		Name:          types.String{Value: connection.Name},
		Schedule:      schedule,
		Status:        connection.Status,
	}

	diags = response.State.Set(ctx, result)
	response.Diagnostics.Append(diags...)

	if response.Diagnostics.HasError() {
		return
	}
}

func (r resourceConnection) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {
	// Get current state
	var state Connection
	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	sourceId := state.ConnectionId.Value

	connnection, _, err := r.p.client.api.ConnectionApi.GetConnection(ctx).ConnectionIdRequestBody(airbyte_sdk.ConnectionIdRequestBody{
		ConnectionId: sourceId,
	}).Execute()

	if err != nil {
		response.Diagnostics.AddError(
			"Error reading source",
			"Could not read sourceID "+sourceId+": "+err.Error(),
		)
		return
	}

	state.Name = types.String{Value: connnection.Name}

	// Set state
	diags = response.State.Set(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (r resourceConnection) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	panic("implement me")
}

func (r resourceConnection) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	var state Connection
	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	// Get source ID from state
	connectionId := state.ConnectionId.Value

	// Delete source by calling API
	_, err := r.p.client.api.ConnectionApi.DeleteConnection(ctx).ConnectionIdRequestBody(airbyte_sdk.ConnectionIdRequestBody{ConnectionId: connectionId}).Execute()
	if err != nil {
		response.Diagnostics.AddError(
			"Error deleting connection",
			"Could not delete connectionId "+connectionId+": "+err.Error(),
		)
		return
	}

	// Remove resource from state
	response.State.RemoveResource(ctx)
}

func (r resourceConnection) ImportState(ctx context.Context, request tfsdk.ImportResourceStateRequest, response *tfsdk.ImportResourceStateResponse) {
	panic("implement me")
}
