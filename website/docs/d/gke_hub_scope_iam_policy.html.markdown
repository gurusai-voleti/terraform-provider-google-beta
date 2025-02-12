---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This file is automatically generated by Magic Modules and manual
#     changes will be clobbered when the file is regenerated.
#
#     Please read more about how to change this file in
#     .github/CONTRIBUTING.md.
#
# ----------------------------------------------------------------------------
subcategory: "GKEHub"
description: |-
  A datasource to retrieve the IAM policy state for GKEHub Scope
---


# google_gke_hub_scope_iam_policy

Retrieves the current IAM policy data for scope


## Example Usage


```hcl
data "google_gke_hub_scope_iam_policy" "policy" {
  project = google_gke_hub_scope.scope.project
  scope_id = google_gke_hub_scope.scope.scope_id
}
```

## Argument Reference

The following arguments are supported:


* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.

## Attributes Reference

The attributes are exported:

* `etag` - (Computed) The etag of the IAM policy.

* `policy_data` - (Required only by `google_gke_hub_scope_iam_policy`) The policy data generated by
  a `google_iam_policy` data source.
