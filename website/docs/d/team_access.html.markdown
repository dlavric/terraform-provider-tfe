---
layout: "tfe"
page_title: "Terraform Enterprise: tfe_team_access"
sidebar_current: "docs-datasource-tfe-team-access"
description: |-
  Get information on team permissions on a workspace.
---

# Data Source: tfe_team_access

Use this data source to get information about team permissions for a workspace.

## Example Usage

```hcl
data "tfe_team_access" "test" {
  team_id      = "my-team-id"
  workspace_id = "my-workspace-id"
}
```

## Argument Reference

The following arguments are supported:

* `team_id` - (Required) ID of the team.
* `workspace_id` - (Required) ID of the workspace.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` The team access ID.
* `access` - The type of access granted to the team on the workspace.
* `permissions` - The permissions granted to the team on the workspaces for each whatever.

The `permissions` block contains:

* `runs` - The permission granted to runs. Valid values are `read`, `plan`, or `apply`
* `variables` - The permissions granted to variables. Valid values are `none`, `read`, or `write`
* `state_versions` - The permissions granted to state versions. Valid values are `none`, `read-outputs`, `read`, or `write`
* `sentinel_mocks` - The permissions granted to Sentinel mocks. Valid values are `none` or `read`
* `workspace_locking` - Whether permission is granted to manually lock the workspace or not.
* `run_tasks` - Boolean determining whether or not to grant the team permission to manage workspace run tasks.
