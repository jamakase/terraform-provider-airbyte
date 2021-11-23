package airbyte

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jamakase/terraform-prodiver-airbyte/airbyte_sdk"
)

// Workspace -
type Workspace struct {
	WorkspaceId types.String `tfsdk:"workspace_id"`
	Name        types.String `tfsdk:"name"`
	CustomerId  types.String `tfsdk:"customer_id"`
	Email       types.String `tfsdk:"email"`
	//Slug                    string          `json:"slug"`
	//InitialSetupComplete    bool            `json:"initialSetupComplete"`
	//DisplaySetupWizard      *bool           `json:"displaySetupWizard,omitempty"`
	//AnonymousDataCollection *bool           `json:"anonymousDataCollection,omitempty"`
	//News                    *bool           `json:"news,omitempty"`
	//SecurityUpdates         *bool           `json:"securityUpdates,omitempty"`
	//Notifications           *[]Notification `json:"notifications,omitempty"`
}

type Source struct {
	SourceDefinitionId types.String `tfsdk:"source_definition_id"`
	SourceId           types.String `tfsdk:"source_id"`
	WorkspaceId        types.String `tfsdk:"workspace_id"`
	// The values required to configure the source. The schema for this must match the schema return by source_definition_specifications/get for the sourceDefinition.
	ConnectionConfiguration types.String `tfsdk:"connection_configuration"`
	Name                    types.String `tfsdk:"name"`
}

type Destination struct {
	DestinationDefinitionId types.String `tfsdk:"destination_definition_id"`
	DestinationId           types.String `tfsdk:"destination_id"`
	WorkspaceId             types.String `tfsdk:"workspace_id"`
	// The values required to configure the destination. The schema for this must match the schema return by destination_definition_specifications/get for the destinationDefinition.
	ConnectionConfiguration types.String `tfsdk:"connection_configuration"`
	Name                    types.String `tfsdk:"name"`
}

type Connection struct {
	ConnectionId types.String `tfsdk:"connection_id"`
	Name         types.String `tfsdk:"name"`
	//NamespaceDefinition  types.Object                 `tfsdk:"namespace_definition"`
	//NamespaceFormat      types.String                 `tfsdk:"namespace_format"`
	//Prefix               types.String                 `tfsdk:"prefix"`
	SourceId      types.String `tfsdk:"source_id"`
	DestinationId types.String `tfsdk:"destination_id"`
	//OperationIds         types.List                   `tfsdk:"operation_ids"`
	//SyncCatalog          types.Object                 `tfsdk:"sync_catalog"`
	Schedule *Schedule                    `tfsdk:"schedule"`
	Status   airbyte_sdk.ConnectionStatus `tfsdk:"status"`
	//ResourceRequirements types.Object                 `tfsdk:"resource_requirements"`
}

type Schedule struct {
	Units    types.Int64  `tfsdk:"units"`
	TimeUnit types.String `tfsdk:"time_unit"`
}

type StreamAndConfiguration struct {
	Stream types.Object `tfsdk:"stream"`
	Config types.Object `tfsdk:"config"`
}

type Stream struct {
	// Stream's name.
	Name string `json:"name"`
}

type StreamFullRefreshConfig struct {
	// Stream's name.
	DestinationSyncMode types.String `tfsdk:"destinationSyncMode"`
	// Alias name to the stream to be used in the destination
	AliasName types.String `tfsdk:"aliasName"`
	Selected  types.Bool   `tfsdk:"selected"`
}

type StreamIncrementalConfig struct {
	// Path to the field that will be used to determine if a record is new or modified since the last sync. This field is REQUIRED if `sync_mode` is `incremental`. Otherwise it is ignored.
	CursorField         *[]string    `tfsdk:"cursorField"`
	DestinationSyncMode types.String `tfsdk:"destinationSyncMode"`
	// Paths to the fields that will be used as primary key. This field is REQUIRED if `destination_sync_mode` is `*_dedup`. Otherwise it is ignored.
	PrimaryKey *[][]string `tfsdk:"primaryKey"`
	// Alias name to the stream to be used in the destination
	AliasName *string `tfsdk:"aliasName"`
	Selected  *bool   `tfsdk:"selected"`
}

type SourceDefinition struct {
	SourceDefinitionId types.String `tfsdk:"source_definition_id"`
	Name               types.String `tfsdk:"name"`
	DockerRepository   types.String `tfsdk:"docker_repository"`
	DockerImageTag     types.String `tfsdk:"docker_image_tag"`
	DocumentationUrl   types.String `tfsdk:"documentation_url"`
	Icon               types.String `tfsdk:"icon"`
}

type DestinationDefinition struct {
	DestinationDefinitionId types.String `tfsdk:"destination_definition_id"`
	Name                    types.String `tfsdk:"name"`
	DockerRepository        types.String `tfsdk:"docker_repository"`
	DockerImageTag          types.String `tfsdk:"docker_image_tag"`
	DocumentationUrl        types.String `tfsdk:"documentation_url"`
	Icon                    types.String `tfsdk:"icon"`
}
