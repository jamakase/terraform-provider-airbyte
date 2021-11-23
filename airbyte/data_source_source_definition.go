package airbyte

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type dataSourceSourceDefinitionType struct{}

func (r dataSourceSourceDefinitionType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"source_definition_id": {
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

func (r dataSourceSourceDefinitionType) NewDataSource(ctx context.Context, p tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return dataSourceSourceDefinition{
		p: *(p.(*provider)),
	}, nil
}

type dataSourceSourceDefinition struct {
	p provider
}

func (r dataSourceSourceDefinition) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {

	var plan SourceDefinition
	diags := req.Config.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	srcDefList, _, err := r.p.client.api.SourceDefinitionApi.ListSourceDefinitions(ctx).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error retrieving source definitions",
			err.Error(),
		)
		return
	}

	// Declare struct that this function will set to this data source's state
	var resourceState *SourceDefinition

	for _, definition := range srcDefList.SourceDefinitions {
		if definition.Name == plan.Name.Value {
			resourceState = &SourceDefinition{
				SourceDefinitionId: types.String{
					Value: definition.SourceDefinitionId,
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
