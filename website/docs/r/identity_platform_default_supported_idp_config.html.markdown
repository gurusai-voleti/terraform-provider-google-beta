---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/identityplatform/DefaultSupportedIdpConfig.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Identity Platform"
description: |-
  Configurations options for authenticating with a the standard set of Identity Toolkit-trusted IDPs.
---

# google_identity_platform_default_supported_idp_config

Configurations options for authenticating with a the standard set of Identity Toolkit-trusted IDPs.

You must enable the
[Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
the marketplace prior to using this resource.



## Example Usage - Identity Platform Default Supported Idp Config Basic


```hcl
resource "google_identity_platform_default_supported_idp_config" "idp_config" {
  enabled       = true
  idp_id        = "playgames.google.com"
  client_id     = "client-id"
  client_secret = "secret"
}
```

## Argument Reference

The following arguments are supported:


* `idp_id` -
  (Required)
  ID of the IDP. Possible values include:
  * `apple.com`
  * `facebook.com`
  * `gc.apple.com`
  * `github.com`
  * `google.com`
  * `linkedin.com`
  * `microsoft.com`
  * `playgames.google.com`
  * `twitter.com`
  * `yahoo.com`

* `client_id` -
  (Required)
  OAuth client ID

* `client_secret` -
  (Required)
  OAuth client secret


* `enabled` -
  (Optional)
  If this IDP allows the user to sign in

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}`

* `name` -
  The name of the DefaultSupportedIdpConfig resource


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


DefaultSupportedIdpConfig can be imported using any of these accepted formats:

* `projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}`
* `{{project}}/{{idp_id}}`
* `{{idp_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import DefaultSupportedIdpConfig using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}"
  to = google_identity_platform_default_supported_idp_config.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), DefaultSupportedIdpConfig can be imported using one of the formats above. For example:

```
$ terraform import google_identity_platform_default_supported_idp_config.default projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}
$ terraform import google_identity_platform_default_supported_idp_config.default {{project}}/{{idp_id}}
$ terraform import google_identity_platform_default_supported_idp_config.default {{idp_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
