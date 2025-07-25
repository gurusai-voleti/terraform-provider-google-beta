---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/edgecontainer/Cluster.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Google Distributed Cloud Edge"
description: |-
  Cluster contains information about a Google Distributed Cloud Edge Kubernetes cluster.
---

# google_edgecontainer_cluster

Cluster contains information about a Google Distributed Cloud Edge Kubernetes cluster.


To get more information about Cluster, see:

* [API documentation](https://cloud.google.com/distributed-cloud/edge/latest/docs/reference/container/rest/v1/projects.locations.clusters)
* How-to Guides
    * [Create and manage clusters](https://cloud.google.com/distributed-cloud/edge/latest/docs/clusters)

~> **Warning:** All arguments including the following potentially sensitive
values will be stored in the raw state as plain text: `cluster_ca_certificate`.
[Read more about sensitive data in state](https://www.terraform.io/language/state/sensitive-data).

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=edgecontainer_cluster&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Edgecontainer Cluster


```hcl
resource "google_edgecontainer_cluster" "default" {
  name = "basic-cluster"
  location = "us-central1"

  authorization {
    admin_users {
      username = "admin@hashicorptest.com"
    }
  }

  networking {
    cluster_ipv4_cidr_blocks = ["10.0.0.0/16"]
    services_ipv4_cidr_blocks = ["10.1.0.0/16"]
  }

  fleet {
    project = "projects/${data.google_project.project.number}"
  }

  labels = {
    my_key    = "my_val"
    other_key = "other_val"
  }
}

data "google_project" "project" {}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=edgecontainer_cluster_with_maintenance_window&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Edgecontainer Cluster With Maintenance Window


```hcl
resource "google_edgecontainer_cluster" "default" {
  name = "cluster-with-maintenance"
  location = "us-central1"

  authorization {
    admin_users {
      username = "admin@hashicorptest.com"
    }
  }

  networking {
    cluster_ipv4_cidr_blocks = ["10.0.0.0/16"]
    services_ipv4_cidr_blocks = ["10.1.0.0/16"]
  }

  fleet {
    project = "projects/${data.google_project.project.number}"
  }

  maintenance_policy {
    window {
      recurring_window {
        window {
          start_time = "2023-01-01T08:00:00Z"
          end_time = "2023-01-01T17:00:00Z"
        }

        recurrence = "FREQ=WEEKLY;BYDAY=SA"
      }
    }
  }
}

data "google_project" "project" {}
```
## Example Usage - Edgecontainer Local Control Plane Cluster


```hcl
resource "google_edgecontainer_cluster" "default" {
  name = "local-control-plane-cluster"
  location = "us-central1"

  authorization {
    admin_users {
      username = "admin@hashicorptest.com"
    }
  }

  networking {
    cluster_ipv4_cidr_blocks = ["10.0.0.0/16"]
    services_ipv4_cidr_blocks = ["10.1.0.0/16"]
  }

  fleet {
    project = "projects/${data.google_project.project.number}"
  }

  external_load_balancer_ipv4_address_pools = ["10.100.0.0-10.100.0.10"]

  control_plane {
    local {
      node_location = "us-central1-edge-example-edgesite"
      node_count = 1
      machine_filter = "machine-name"
      shared_deployment_policy = "ALLOWED"
    }
  }
}

data "google_project" "project" {}
```

## Argument Reference

The following arguments are supported:


* `fleet` -
  (Required)
  Fleet related configuration.
  Fleets are a Google Cloud concept for logically organizing clusters,
  letting you use and manage multi-cluster capabilities and apply
  consistent policies across your systems.
  Structure is [documented below](#nested_fleet).

* `networking` -
  (Required)
  Fleet related configuration.
  Fleets are a Google Cloud concept for logically organizing clusters,
  letting you use and manage multi-cluster capabilities and apply
  consistent policies across your systems.
  Structure is [documented below](#nested_networking).

* `authorization` -
  (Required)
  RBAC policy that will be applied and managed by GEC.
  Structure is [documented below](#nested_authorization).

* `location` -
  (Required)
  The location of the resource.

* `name` -
  (Required)
  The GDCE cluster name.


* `labels` -
  (Optional)
  User-defined labels for the edgecloud cluster.
  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `default_max_pods_per_node` -
  (Optional)
  The default maximum number of pods per node used if a maximum value is not
  specified explicitly for a node pool in this cluster. If unspecified, the
  Kubernetes default value will be used.

* `maintenance_policy` -
  (Optional)
  Cluster-wide maintenance policy configuration.
  Structure is [documented below](#nested_maintenance_policy).

* `control_plane` -
  (Optional)
  The configuration of the cluster control plane.
  Structure is [documented below](#nested_control_plane).

* `system_addons_config` -
  (Optional)
  Config that customers are allowed to define for GDCE system add-ons.
  Structure is [documented below](#nested_system_addons_config).

* `external_load_balancer_ipv4_address_pools` -
  (Optional)
  Address pools for cluster data plane external load balancing.

* `control_plane_encryption` -
  (Optional)
  Remote control plane disk encryption options. This field is only used when
  enabling CMEK support.
  Structure is [documented below](#nested_control_plane_encryption).

* `target_version` -
  (Optional)
  The target cluster version. For example: "1.5.0".

* `release_channel` -
  (Optional)
  The release channel a cluster is subscribed to.
  Possible values are: `RELEASE_CHANNEL_UNSPECIFIED`, `NONE`, `REGULAR`.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_fleet"></a>The `fleet` block supports:

* `project` -
  (Required)
  The name of the Fleet host project where this cluster will be registered.
  Project names are formatted as
  `projects/<project-number>`.

* `membership` -
  (Output)
  The name of the managed Hub Membership resource associated to this cluster.
  Membership names are formatted as
  `projects/<project-number>/locations/global/membership/<cluster-id>`.

<a name="nested_networking"></a>The `networking` block supports:

* `cluster_ipv4_cidr_blocks` -
  (Required)
  All pods in the cluster are assigned an RFC1918 IPv4 address from these
  blocks. Only a single block is supported. This field cannot be changed
  after creation.

* `services_ipv4_cidr_blocks` -
  (Required)
  All services in the cluster are assigned an RFC1918 IPv4 address from these
  blocks. Only a single block is supported. This field cannot be changed
  after creation.

* `cluster_ipv6_cidr_blocks` -
  (Optional)
  If specified, dual stack mode is enabled and all pods in the cluster are
  assigned an IPv6 address from these blocks alongside from an IPv4
  address. Only a single block is supported. This field cannot be changed
  after creation.

* `services_ipv6_cidr_blocks` -
  (Optional)
  If specified, dual stack mode is enabled and all services in the cluster are
  assigned an IPv6 address from these blocks alongside from an IPv4
  address. Only a single block is supported. This field cannot be changed
  after creation.

* `network_type` -
  (Output)
  IP addressing type of this cluster i.e. SINGLESTACK_V4 vs DUALSTACK_V4_V6.

<a name="nested_authorization"></a>The `authorization` block supports:

* `admin_users` -
  (Required)
  User that will be granted the cluster-admin role on the cluster, providing
  full access to the cluster. Currently, this is a singular field, but will
  be expanded to allow multiple admins in the future.
  Structure is [documented below](#nested_authorization_admin_users).


<a name="nested_authorization_admin_users"></a>The `admin_users` block supports:

* `username` -
  (Required)
  An active Google username.

<a name="nested_maintenance_policy"></a>The `maintenance_policy` block supports:

* `window` -
  (Required)
  Specifies the maintenance window in which maintenance may be performed.
  Structure is [documented below](#nested_maintenance_policy_window).

* `maintenance_exclusions` -
  (Optional)
  Exclusions to automatic maintenance. Non-emergency maintenance should not occur
  in these windows. Each exclusion has a unique name and may be active or expired.
  The max number of maintenance exclusions allowed at a given time is 3.
  Structure is [documented below](#nested_maintenance_policy_maintenance_exclusions).


<a name="nested_maintenance_policy_window"></a>The `window` block supports:

* `recurring_window` -
  (Required)
  Represents an arbitrary window of time that recurs.
  Structure is [documented below](#nested_maintenance_policy_window_recurring_window).


<a name="nested_maintenance_policy_window_recurring_window"></a>The `recurring_window` block supports:

* `window` -
  (Optional)
  Represents an arbitrary window of time.
  Structure is [documented below](#nested_maintenance_policy_window_recurring_window_window).

* `recurrence` -
  (Optional)
  An RRULE (https://tools.ietf.org/html/rfc5545#section-3.8.5.3) for how
  this window recurs. They go on for the span of time between the start and
  end time.


<a name="nested_maintenance_policy_window_recurring_window_window"></a>The `window` block supports:

* `start_time` -
  (Optional)
  The time that the window first starts.

* `end_time` -
  (Optional)
  The time that the window ends. The end time must take place after the
  start time.

<a name="nested_maintenance_policy_maintenance_exclusions"></a>The `maintenance_exclusions` block supports:

* `window` -
  (Optional)
  Represents an arbitrary window of time.
  Structure is [documented below](#nested_maintenance_policy_maintenance_exclusions_maintenance_exclusions_window).

* `id` -
  (Optional)
  A unique (per cluster) id for the window.


<a name="nested_maintenance_policy_maintenance_exclusions_maintenance_exclusions_window"></a>The `window` block supports:

* `start_time` -
  (Optional)
  The time that the window first starts.

* `end_time` -
  (Optional)
  The time that the window ends. The end time must take place after the
  start time.

<a name="nested_control_plane"></a>The `control_plane` block supports:

* `remote` -
  (Optional)
  Remote control plane configuration.
  Structure is [documented below](#nested_control_plane_remote).

* `local` -
  (Optional)
  Local control plane configuration.
  Structure is [documented below](#nested_control_plane_local).


<a name="nested_control_plane_remote"></a>The `remote` block supports:

* `node_location` -
  (Optional)
  Name of the Google Distributed Cloud Edge zones where this node pool
  will be created. For example: `us-central1-edge-customer-a`.

<a name="nested_control_plane_local"></a>The `local` block supports:

* `node_location` -
  (Optional)
  Name of the Google Distributed Cloud Edge zones where this node pool
  will be created. For example: `us-central1-edge-customer-a`.

* `node_count` -
  (Optional)
  The number of nodes to serve as replicas of the Control Plane.
  Only 1 and 3 are supported.

* `machine_filter` -
  (Optional)
  Only machines matching this filter will be allowed to host control
  plane nodes. The filtering language accepts strings like "name=<name>",
  and is documented here: [AIP-160](https://google.aip.dev/160).

* `shared_deployment_policy` -
  (Optional)
  Policy configuration about how user applications are deployed.
  Possible values are: `SHARED_DEPLOYMENT_POLICY_UNSPECIFIED`, `ALLOWED`, `DISALLOWED`.

<a name="nested_system_addons_config"></a>The `system_addons_config` block supports:

* `ingress` -
  (Optional)
  Config for the Ingress add-on which allows customers to create an Ingress
  object to manage external access to the servers in a cluster. The add-on
  consists of istiod and istio-ingress.
  Structure is [documented below](#nested_system_addons_config_ingress).


<a name="nested_system_addons_config_ingress"></a>The `ingress` block supports:

* `disabled` -
  (Optional)
  Whether Ingress is disabled.

* `ipv4_vip` -
  (Optional)
  Ingress VIP.

<a name="nested_control_plane_encryption"></a>The `control_plane_encryption` block supports:

* `kms_key` -
  (Optional)
  The Cloud KMS CryptoKey e.g.
  projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
  to use for protecting control plane disks. If not specified, a
  Google-managed key will be used instead.

* `kms_key_active_version` -
  (Output)
  The Cloud KMS CryptoKeyVersion currently in use for protecting control
  plane disks. Only applicable if kms_key is set.

* `kms_key_state` -
  (Output)
  Availability of the Cloud KMS CryptoKey. If not `KEY_AVAILABLE`, then
  nodes may go offline as they cannot access their local data. This can be
  caused by a lack of permissions to use the key, or if the key is disabled
  or deleted.

* `kms_status` -
  (Output)
  Error status returned by Cloud KMS when using this key. This field may be
  populated only if `kms_key_state` is not `KMS_KEY_STATE_KEY_AVAILABLE`.
  If populated, this field contains the error status reported by Cloud KMS.
  Structure is [documented below](#nested_control_plane_encryption_kms_status).


<a name="nested_control_plane_encryption_kms_status"></a>The `kms_status` block contains:

* `code` -
  (Output)
  The status code, which should be an enum value of google.rpc.Code.

* `message` -
  (Output)
  A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/clusters/{{name}}`

* `create_time` -
  The time the cluster was created, in RFC3339 text format.

* `update_time` -
  The time the cluster was last updated, in RFC3339 text format.

* `endpoint` -
  The IP address of the Kubernetes API server.

* `port` -
  The port number of the Kubernetes API server.

* `cluster_ca_certificate` -
  The PEM-encoded public certificate of the cluster's CA.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `control_plane_version` -
  The control plane release version.

* `node_version` -
  The lowest release version among all worker nodes. This field can be empty
  if the cluster does not have any worker nodes.

* `status` -
  Indicates the status of the cluster.

* `maintenance_events` -
  All the maintenance events scheduled for the cluster, including the ones
  ongoing, planned for the future and done in the past (up to 90 days).
  Structure is [documented below](#nested_maintenance_events).

* `terraform_labels` -
  The combination of labels configured directly on the resource
   and default labels configured on the provider.

* `effective_labels` -
  All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.


<a name="nested_maintenance_events"></a>The `maintenance_events` block contains:

* `uuid` -
  (Output)
  UUID of the maintenance event.

* `target_version` -
  (Output)
  The target version of the cluster.

* `operation` -
  (Output)
  The operation for running the maintenance event. Specified in the format
  projects/*/locations/*/operations/*. If the maintenance event is split
  into multiple operations (e.g. due to maintenance windows), the latest
  one is recorded.

* `type` -
  (Output)
  Indicates the maintenance event type.

* `schedule` -
  (Output)
  The schedule of the maintenance event.

* `state` -
  (Output)
  Indicates the maintenance event state.

* `create_time` -
  (Output)
  The time when the maintenance event request was created.

* `start_time` -
  (Output)
  The time when the maintenance event started.

* `end_time` -
  (Output)
  The time when the maintenance event ended, either successfully or not. If
  the maintenance event is split into multiple maintenance windows,
  end_time is only updated when the whole flow ends.

* `update_time` -
  (Output)
  The time when the maintenance event message was updated.

## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 480 minutes.
- `update` - Default is 480 minutes.
- `delete` - Default is 480 minutes.

## Import


Cluster can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/clusters/{{name}}`
* `{{project}}/{{location}}/{{name}}`
* `{{location}}/{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Cluster using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/clusters/{{name}}"
  to = google_edgecontainer_cluster.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Cluster can be imported using one of the formats above. For example:

```
$ terraform import google_edgecontainer_cluster.default projects/{{project}}/locations/{{location}}/clusters/{{name}}
$ terraform import google_edgecontainer_cluster.default {{project}}/{{location}}/{{name}}
$ terraform import google_edgecontainer_cluster.default {{location}}/{{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
