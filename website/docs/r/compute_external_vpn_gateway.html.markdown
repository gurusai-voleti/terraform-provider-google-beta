---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/ExternalVpnGateway.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Compute Engine"
description: |-
  Represents a VPN gateway managed outside of GCP.
---

# google_compute_external_vpn_gateway

Represents a VPN gateway managed outside of GCP.


To get more information about ExternalVpnGateway, see:

* [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/externalVpnGateways)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=external_vpn_gateway&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - External Vpn Gateway


```hcl
resource "google_compute_ha_vpn_gateway" "ha_gateway" {
  region   = "us-central1"
  name     = "ha-vpn"
  network  = google_compute_network.network.id
}

resource "google_compute_external_vpn_gateway" "external_gateway" {
  name            = "external-gateway"
  redundancy_type = "SINGLE_IP_INTERNALLY_REDUNDANT"
  description     = "An externally managed VPN gateway"
  interface {
    id         = 0
    ip_address = "8.8.8.8"
  }
}

resource "google_compute_network" "network" {
  name                    = "network-1"
  routing_mode            = "GLOBAL"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "network_subnet1" {
  name          = "ha-vpn-subnet-1"
  ip_cidr_range = "10.0.1.0/24"
  region        = "us-central1"
  network       = google_compute_network.network.id
}

resource "google_compute_subnetwork" "network_subnet2" {
  name          = "ha-vpn-subnet-2"
  ip_cidr_range = "10.0.2.0/24"
  region        = "us-west1"
  network       = google_compute_network.network.id
}

resource "google_compute_router" "router1" {
  name     = "ha-vpn-router1"
  network  = google_compute_network.network.name
  bgp {
    asn = 64514
  }
}

resource "google_compute_vpn_tunnel" "tunnel1" {
  name                            = "ha-vpn-tunnel1"
  region                          = "us-central1"
  vpn_gateway                     = google_compute_ha_vpn_gateway.ha_gateway.id
  peer_external_gateway           = google_compute_external_vpn_gateway.external_gateway.id
  peer_external_gateway_interface = 0
  shared_secret                   = "a secret message"
  router                          = google_compute_router.router1.id
  vpn_gateway_interface           = 0
}

resource "google_compute_vpn_tunnel" "tunnel2" {
  name                            = "ha-vpn-tunnel2"
  region                          = "us-central1"
  vpn_gateway                     = google_compute_ha_vpn_gateway.ha_gateway.id
  peer_external_gateway           = google_compute_external_vpn_gateway.external_gateway.id
  peer_external_gateway_interface = 0
  shared_secret                   = "a secret message"
  router                          = " ${google_compute_router.router1.id}"
  vpn_gateway_interface           = 1
}

resource "google_compute_router_interface" "router1_interface1" {
  name       = "router1-interface1"
  router     = google_compute_router.router1.name
  region     = "us-central1"
  ip_range   = "169.254.0.1/30"
  vpn_tunnel = google_compute_vpn_tunnel.tunnel1.name
}

resource "google_compute_router_peer" "router1_peer1" {
  name                      = "router1-peer1"
  router                    = google_compute_router.router1.name
  region                    = "us-central1"
  peer_ip_address           = "169.254.0.2"
  peer_asn                  = 64515
  advertised_route_priority = 100
  interface                 = google_compute_router_interface.router1_interface1.name
}

resource "google_compute_router_interface" "router1_interface2" {
  name       = "router1-interface2"
  router     = google_compute_router.router1.name
  region     = "us-central1"
  ip_range   = "169.254.1.1/30"
  vpn_tunnel = google_compute_vpn_tunnel.tunnel2.name
}

resource "google_compute_router_peer" "router1_peer2" {
  name                      = "router1-peer2"
  router                    = google_compute_router.router1.name
  region                    = "us-central1"
  peer_ip_address           = "169.254.1.2"
  peer_asn                  = 64515
  advertised_route_priority = 100
  interface                 = google_compute_router_interface.router1_interface2.name
}
```

## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  Name of the resource. Provided by the client when the resource is
  created. The name must be 1-63 characters long, and comply with
  RFC1035.  Specifically, the name must be 1-63 characters long and
  match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means
  the first character must be a lowercase letter, and all following
  characters must be a dash, lowercase letter, or digit, except the last
  character, which cannot be a dash.


* `description` -
  (Optional)
  An optional description of this resource.

* `labels` -
  (Optional)
  Labels for the external VPN gateway resource.
  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `redundancy_type` -
  (Optional)
  Indicates the redundancy type of this external VPN gateway
  Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.

* `interface` -
  (Optional)
  A list of interfaces on this external VPN gateway.
  Structure is [documented below](#nested_interface).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_interface"></a>The `interface` block supports:

* `id` -
  (Optional)
  The numeric ID for this interface. Allowed values are based on the redundancy type
  of this external VPN gateway
  * `0 - SINGLE_IP_INTERNALLY_REDUNDANT`
  * `0, 1 - TWO_IPS_REDUNDANCY`
  * `0, 1, 2, 3 - FOUR_IPS_REDUNDANCY`

* `ip_address` -
  (Optional)
  IP address of the interface in the external VPN gateway.
  Only IPv4 is supported. This IP address can be either from
  your on-premise gateway or another Cloud provider's VPN gateway,
  it cannot be an IP address from Google Compute Engine.

* `ipv6_address` -
  (Optional)
  IPv6 address of the interface in the external VPN gateway. This IPv6
  address can be either from your on-premise gateway or another Cloud
  provider's VPN gateway, it cannot be an IP address from Google Compute
  Engine. Must specify an IPv6 address (not IPV4-mapped) using any format
  described in RFC 4291 (e.g. 2001:db8:0:0:2d9:51:0:0). The output format
  is RFC 5952 format (e.g. 2001:db8::2d9:51:0:0).

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/global/externalVpnGateways/{{name}}`

* `label_fingerprint` -
  The fingerprint used for optimistic locking of this resource.  Used
  internally during updates.

* `terraform_labels` -
  The combination of labels configured directly on the resource
   and default labels configured on the provider.

* `effective_labels` -
  All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.
* `self_link` - The URI of the created resource.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


ExternalVpnGateway can be imported using any of these accepted formats:

* `projects/{{project}}/global/externalVpnGateways/{{name}}`
* `{{project}}/{{name}}`
* `{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ExternalVpnGateway using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/global/externalVpnGateways/{{name}}"
  to = google_compute_external_vpn_gateway.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), ExternalVpnGateway can be imported using one of the formats above. For example:

```
$ terraform import google_compute_external_vpn_gateway.default projects/{{project}}/global/externalVpnGateways/{{name}}
$ terraform import google_compute_external_vpn_gateway.default {{project}}/{{name}}
$ terraform import google_compute_external_vpn_gateway.default {{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
