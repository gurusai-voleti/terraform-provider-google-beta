---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networksecurity/InterceptDeploymentGroup.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Network Security"
description: |-
  A deployment group aggregates many zonal intercept backends (deployments)
  into a single global intercept service.
---

# google_network_security_intercept_deployment_group

A deployment group aggregates many zonal intercept backends (deployments)
into a single global intercept service. Consumers can connect this service
using an endpoint group.



<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=network_security_intercept_deployment_group_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Network Security Intercept Deployment Group Basic


```hcl
resource "google_compute_network" "network" {
  name                    = "example-network"
  auto_create_subnetworks = false
}

resource "google_network_security_intercept_deployment_group" "default" {
  intercept_deployment_group_id = "example-dg"
  location                      = "global"
  network                       = google_compute_network.network.id
  description                   = "some description"
  labels = {
    foo = "bar"
  }
}
```

## Argument Reference

The following arguments are supported:


* `network` -
  (Required)
  The network that will be used for all child deployments, for example:
  `projects/{project}/global/networks/{network}`.
  See https://google.aip.dev/124.

* `location` -
  (Required)
  The cloud location of the deployment group, currently restricted to `global`.

* `intercept_deployment_group_id` -
  (Required)
  The ID to use for the new deployment group, which will become the final
  component of the deployment group's resource name.


* `labels` -
  (Optional)
  Labels are key/value pairs that help to organize and filter resources.
  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `description` -
  (Optional)
  User-provided description of the deployment group.
  Used as additional context for the deployment group.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/interceptDeploymentGroups/{{intercept_deployment_group_id}}`

* `name` -
  The resource name of this deployment group, for example:
  `projects/123456789/locations/global/interceptDeploymentGroups/my-dg`.
  See https://google.aip.dev/122 for more details.

* `create_time` -
  The timestamp when the resource was created.
  See https://google.aip.dev/148#timestamps.

* `update_time` -
  The timestamp when the resource was most recently updated.
  See https://google.aip.dev/148#timestamps.

* `connected_endpoint_groups` -
  The list of endpoint groups that are connected to this resource.
  Structure is [documented below](#nested_connected_endpoint_groups).

* `state` -
  The current state of the deployment group.
  See https://google.aip.dev/216.
  Possible values:
  STATE_UNSPECIFIED
  ACTIVE
  CREATING
  DELETING

* `reconciling` -
  The current state of the resource does not match the user's intended state,
  and the system is working to reconcile them. This is part of the normal
  operation (e.g. adding a new deployment to the group)
  See https://google.aip.dev/128.

* `locations` -
  The list of locations where the deployment group is present.
  Structure is [documented below](#nested_locations).

* `terraform_labels` -
  The combination of labels configured directly on the resource
   and default labels configured on the provider.

* `effective_labels` -
  All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.


<a name="nested_connected_endpoint_groups"></a>The `connected_endpoint_groups` block contains:

* `name` -
  (Output)
  The connected endpoint group's resource name, for example:
  `projects/123456789/locations/global/interceptEndpointGroups/my-eg`.
  See https://google.aip.dev/124.

<a name="nested_locations"></a>The `locations` block contains:

* `state` -
  (Output)
  The current state of the association in this location.
  Possible values:
  STATE_UNSPECIFIED
  ACTIVE
  OUT_OF_SYNC

* `location` -
  (Output)
  The cloud location, e.g. `us-central1-a` or `asia-south1-b`.

## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


InterceptDeploymentGroup can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/interceptDeploymentGroups/{{intercept_deployment_group_id}}`
* `{{project}}/{{location}}/{{intercept_deployment_group_id}}`
* `{{location}}/{{intercept_deployment_group_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import InterceptDeploymentGroup using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/interceptDeploymentGroups/{{intercept_deployment_group_id}}"
  to = google_network_security_intercept_deployment_group.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), InterceptDeploymentGroup can be imported using one of the formats above. For example:

```
$ terraform import google_network_security_intercept_deployment_group.default projects/{{project}}/locations/{{location}}/interceptDeploymentGroups/{{intercept_deployment_group_id}}
$ terraform import google_network_security_intercept_deployment_group.default {{project}}/{{location}}/{{intercept_deployment_group_id}}
$ terraform import google_network_security_intercept_deployment_group.default {{location}}/{{intercept_deployment_group_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
