resource "airbyte_workspace" "default_workspace" {
  name  = "new-workspace"
  email = "test@jamakase.com"
  initialSetup = true
  #  news = "new-workspace"
  #  securityUpdates = "new-workspace"
  #  notifications = "new-workspace"
}

data "airbyte_source_definition" "postgres" {
  name = "Postgres"
}

data "airbyte_source_definition" "pokemon" {
  name = "PokeAPI"
}

data "airbyte_destination_definition" "local" {
  name = "Local JSON"
}

#resource "airbyte_workspace" "default_workspace_2" {
#  name = "workspace-2"
#}

resource "airbyte_source" "source_1" {
  name                 = "poke-source"
  workspace_id         = airbyte_workspace.default_workspace.workspace_id
  source_definition_id = data.airbyte_source_definition.pokemon.source_definition_id

  connection_configuration = jsonencode({
    pokemon_name : "luxray"
  })
}

resource "airbyte_destination" "destination_1" {
  name                      = "destination-2"
  workspace_id              = airbyte_workspace.default_workspace.workspace_id
  destination_definition_id = data.airbyte_destination_definition.local.destination_definition_id

  connection_configuration = jsonencode({
    destination_path : "/local"
  })
}
#
resource "airbyte_connection" "gvcmps_warehouse_pipeline" {
  name           = "GiveCampus app to Warehouse Pipeline"
  source_id      = airbyte_source.source_1.source_id
  destination_id = airbyte_destination.destination_1.destination_id
  status         = "active"
  #  namespace_definition = "destination"
  schedule       = {
    units     = 3
    time_unit = "hours"
  }
  #
  #  incremental_streams = {
  #    destination_sync_mode = "append_dedup"
  #    cursor_field = "updated_at"
  #    primary_key = "id"
  #    names = ["addtnl_fields", "affiliations", "contributions", "designations", ...]
  #  }
  #  incremental_streams = {
  #    destination_sync_mode = "append_dedup"
  #    cursor_field = "other_timestamp"
  #    primary_key = "id"
  #    names = ["subscriptions", "projects", "users", ...]
  #  }
  #  full_refresh_streams = {
  #    destination_sync_mode = "overwrite"
  #    names = "^(campaign_advocate([a-z_]+)|([a-z_]+)forms([a-z_]+)$"
  #  }
}
output "workspace_id" {
  value = airbyte_workspace.default_workspace.workspace_id
}