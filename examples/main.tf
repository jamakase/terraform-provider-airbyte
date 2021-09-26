resource "airbyte_workspace" "default_workspace" {
  name = "new-workspace"
}

resource "airbyte_workspace" "default_workspace_2" {
  name = airbyte_workspace.default_workspace.workspace_id
}