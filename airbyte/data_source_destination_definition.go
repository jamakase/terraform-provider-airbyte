package airbyte

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type dataSourceDestinationDefinitionType struct{}

func (r dataSourceDestinationDefinitionType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"destination_definition_id": {
				Type:     types.StringType,
				Computed: true,
			},
			"name": {
				Type:     types.StringType,
				Computed: true,
			},
			"docker_repository": {
				Type:     types.StringType,
				Computed: true,
			},
			"docker_image_tag": {
				Type:     types.StringType,
				Computed: true,
			},
			"documentation_url": {
				Type:     types.StringType,
				Computed: true,
			},
			"icon": {
				Type:     types.StringType,
				Computed: true,
			},
		},
	}, nil
}

func (r dataSourceDestinationDefinitionType) NewDataSource(ctx context.Context, p tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return dataSourceDestinationDefinition{
		p: *(p.(*provider)),
	}, nil
}

type dataSourceDestinationDefinition struct {
	p provider
}

func (r dataSourceDestinationDefinition) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {

	var plan DestinationDefinition
	diags := req.Config.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	destDefList, _, err := r.p.client.api.DestinationDefinitionApi.ListDestinationDefinitions(ctx).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error retrieving destination definitions",
			err.Error(),
		)
		return
	}

	// Declare struct that this function will set to this data source's state
	var resourceState *DestinationDefinition

	for _, definition := range destDefList.DestinationDefinitions {
		if definition.Name == plan.Name.Value {
			resourceState = &DestinationDefinition{
				DestinationDefinitionId: types.String{
					Value: definition.DestinationDefinitionId,
				},
				Name: types.String{
					Value: definition.Name,
				},
				DockerRepository: types.String{
					Value: definition.DockerRepository,
				},
				DockerImageTag: types.String{
					Value: definition.DockerImageTag,
				},
				DocumentationUrl: types.String{
					//Value: *definition.DocumentationUrl,
					Null: true,
				},
				Icon: types.String{
					//Value: *definition.Icon,
					Null: true,
				},
			}
			break
		}
	}

	if resourceState == nil {
		resp.Diagnostics.AddError(
			"Error retrieving source definitions: not found source definition with name", "not found definition")
		return
	}
	//
	//// Sample debug message
	//// To view this message, set the ATF_LOG environment variable to DEBUG
	//// 		`export TF_LOG=DEBUG`
	//// To hide debug message, unset the environment variable
	//// 		`unset TF_LOG`
	//fmt.Fprintf(os.Stdout, "[DEBUG]-Resource State:%+v", resourceState)

	// Set state
	diags = resp.State.Set(ctx, &resourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
