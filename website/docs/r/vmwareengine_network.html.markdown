---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/vmwareengine/Network.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud VMware Engine"
description: |-
  Provides connectivity for VMware Engine private clouds.
---

# google_vmwareengine_network

Provides connectivity for VMware Engine private clouds.


To get more information about Network, see:

* [API documentation](https://cloud.google.com/vmware-engine/docs/reference/rest/v1/projects.locations.vmwareEngineNetworks)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=vmware_engine_network_standard&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Vmware Engine Network Standard


```hcl
resource "google_vmwareengine_network" "vmw-engine-network" {
    name              = "standard-nw"
    location          = "global" # Standard network needs to be global
    type              = "STANDARD"
    description       = "VMwareEngine standard network sample"
}
```
## Example Usage - Vmware Engine Network Legacy


```hcl
resource "google_vmwareengine_network" "vmw-engine-network" {
  project     = google_project_service.acceptance.project
  name        = "us-west1-default" #Legacy network IDs are in the format: {region-id}-default
  location    = "us-west1"
  type        = "LEGACY"
  description = "VMwareEngine legacy network sample"
}

resource "google_project_service" "acceptance" {
  project  = google_project.acceptance.project_id
  service  = "vmwareengine.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

# there can be only 1 Legacy network per region for a given project,
# so creating new project for isolation in CI.
resource "google_project" "acceptance" {
  name            = "vmw-proj"
  project_id      = "vmw-proj"
  org_id          = "123456789"
  billing_account = "000000-0000000-0000000-000000"
  deletion_policy = "DELETE"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.acceptance]

  create_duration = "60s"
}
```

## Argument Reference

The following arguments are supported:


* `type` -
  (Required)
  VMware Engine network type.
  Possible values are: `LEGACY`, `STANDARD`.

* `location` -
  (Required)
  The location where the VMwareEngineNetwork should reside.

* `name` -
  (Required)
  The ID of the VMwareEngineNetwork.


* `description` -
  (Optional)
  User-provided description for this VMware Engine network.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/vmwareEngineNetworks/{{name}}`

* `vpc_networks` -
  VMware Engine service VPC networks that provide connectivity from a private cloud to customer projects,
  the internet, and other Google Cloud services.
  Structure is [documented below](#nested_vpc_networks).

* `state` -
  State of the VMware Engine network.

* `uid` -
  System-generated unique identifier for the resource.


<a name="nested_vpc_networks"></a>The `vpc_networks` block contains:

* `type` -
  (Output)
  Type of VPC network (INTRANET, INTERNET, or GOOGLE_CLOUD)

* `network` -
  (Output)
  The relative resource name of the service VPC network this VMware Engine network is attached to.
  For example: projects/123123/global/networks/my-network

## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


Network can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/vmwareEngineNetworks/{{name}}`
* `{{project}}/{{location}}/{{name}}`
* `{{location}}/{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Network using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/vmwareEngineNetworks/{{name}}"
  to = google_vmwareengine_network.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Network can be imported using one of the formats above. For example:

```
$ terraform import google_vmwareengine_network.default projects/{{project}}/locations/{{location}}/vmwareEngineNetworks/{{name}}
$ terraform import google_vmwareengine_network.default {{project}}/{{location}}/{{name}}
$ terraform import google_vmwareengine_network.default {{location}}/{{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
