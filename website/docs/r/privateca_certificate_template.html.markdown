---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/privateca/CertificateTemplate.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Certificate Authority Service"
description: |-
  Certificate Authority Service provides reusable and parameterized templates that you can use for common certificate issuance scenarios.
---

# google_privateca_certificate_template

Certificate Authority Service provides reusable and parameterized templates that you can use for common certificate issuance scenarios. A certificate template represents a relatively static and well-defined certificate issuance schema within an organization. A certificate template can essentially become a full-fledged vertical certificate issuance framework.


To get more information about CertificateTemplate, see:

* [API documentation](https://cloud.google.com/certificate-authority-service/docs/reference/rest)
* How-to Guides
    * [Common configurations and Certificate Profiles](https://cloud.google.com/certificate-authority-service/docs/certificate-profile)
    * [Official Documentation](https://cloud.google.com/certificate-authority-service)
    * [Understanding Certificate Templates](https://cloud.google.com/certificate-authority-service/docs/certificate-template)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=privateca_template_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Privateca Template Basic


```hcl
resource "google_privateca_certificate_template" "default" {
  name = "my-template"
  location = "us-central1"
  description = "A sample certificate template"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }

  maximum_lifetime = "86400s"

  passthrough_extensions {
    additional_extensions {
      object_id_path = [1, 6]
    }
    known_extensions = ["EXTENDED_KEY_USAGE"]
  }

  predefined_values {
    additional_extensions {
      object_id {
        object_id_path = [1, 6]
      }
      value    = "c3RyaW5nCg=="
      critical = true
    }
    aia_ocsp_servers = ["string"]
    ca_options {
      is_ca                  = false
      max_issuer_path_length = 6
    }
    key_usage {
      base_key_usage {
        cert_sign          = false
        content_commitment = true
        crl_sign           = false
        data_encipherment  = true
        decipher_only      = true
        digital_signature  = true
        encipher_only      = true
        key_agreement      = true
        key_encipherment   = true
      }
      extended_key_usage {
        client_auth      = true
        code_signing     = true
        email_protection = true
        ocsp_signing     = true
        server_auth      = true
        time_stamping    = true
      }
      unknown_extended_key_usages {
        object_id_path = [1, 6]
      }
    }
    policy_ids {
      object_id_path = [1, 6]
    }
  }

  labels = {
    label-one = "value-one"
  }
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=privateca_template_zero_max_issuer_path_length_null_ca&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Privateca Template Zero Max Issuer Path Length Null Ca


```hcl
resource "google_privateca_certificate_template" "default" {
  name = "my-template"
  location = "us-central1"
  description = "A sample certificate template"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }

  maximum_lifetime = "86400s"

  passthrough_extensions {
    additional_extensions {
      object_id_path = [1, 6]
    }
    known_extensions = ["EXTENDED_KEY_USAGE"]
  }

  predefined_values {
    additional_extensions {
      object_id {
        object_id_path = [1, 6]
      }
      value    = "c3RyaW5nCg=="
      critical = true
    }
    aia_ocsp_servers = ["string"]
    ca_options {
      is_ca                       = false
      null_ca                     = true
      zero_max_issuer_path_length = true
      max_issuer_path_length      = 0
    }
    key_usage {
      base_key_usage {
        cert_sign          = false
        content_commitment = true
        crl_sign           = false
        data_encipherment  = true
        decipher_only      = true
        digital_signature  = true
        encipher_only      = true
        key_agreement      = true
        key_encipherment   = true
      }
      extended_key_usage {
        client_auth      = true
        code_signing     = true
        email_protection = true
        ocsp_signing     = true
        server_auth      = true
        time_stamping    = true
      }
      unknown_extended_key_usages {
        object_id_path = [1, 6]
      }
    }
    policy_ids {
      object_id_path = [1, 6]
    }
    name_constraints {
      critical                  = true
      permitted_dns_names       = ["*.example1.com", "*.example2.com"]
      excluded_dns_names        = ["*.deny.example1.com", "*.deny.example2.com"]
      permitted_ip_ranges       = ["10.0.0.0/8", "11.0.0.0/8"]
      excluded_ip_ranges        = ["10.1.1.0/24", "11.1.1.0/24"]
      permitted_email_addresses = [".example1.com", ".example2.com"]
      excluded_email_addresses  = [".deny.example1.com", ".deny.example2.com"]
      permitted_uris            = [".example1.com", ".example2.com"]
      excluded_uris             = [".deny.example1.com", ".deny.example2.com"]
    }
  }

  labels = {
    label-one = "value-one"
  }
}
```

## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  The resource name for this CertificateTemplate in the format `projects/*/locations/*/certificateTemplates/*`.

* `location` -
  (Required)
  The location for the resource


* `predefined_values` -
  (Optional)
  Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.
  Structure is [documented below](#nested_predefined_values).

* `identity_constraints` -
  (Optional)
  Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
  Structure is [documented below](#nested_identity_constraints).

* `passthrough_extensions` -
  (Optional)
  Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
  Structure is [documented below](#nested_passthrough_extensions).

* `maximum_lifetime` -
  (Optional)
  Optional. The maximum lifetime allowed for all issued certificates that use this template. If the issuing CaPool's IssuancePolicy specifies a maximum lifetime the minimum of the two durations will be the maximum lifetime for issued. Note that if the issuing CertificateAuthority expires before a Certificate's requested maximum_lifetime, the effective lifetime will be explicitly truncated to match it.

* `description` -
  (Optional)
  Optional. A human-readable description of scenarios this template is intended for.

* `labels` -
  (Optional)
  Optional. Labels with user-defined metadata.
  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_predefined_values"></a>The `predefined_values` block supports:

* `key_usage` -
  (Optional)
  Optional. Indicates the intended use for keys that correspond to a certificate.
  Structure is [documented below](#nested_predefined_values_key_usage).

* `ca_options` -
  (Optional)
  Optional. Describes options in this X509Parameters that are relevant in a CA certificate.
  Structure is [documented below](#nested_predefined_values_ca_options).

* `policy_ids` -
  (Optional)
  Optional. Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
  Structure is [documented below](#nested_predefined_values_policy_ids).

* `aia_ocsp_servers` -
  (Optional)
  Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate.

* `additional_extensions` -
  (Optional)
  Optional. Describes custom X.509 extensions.
  Structure is [documented below](#nested_predefined_values_additional_extensions).

* `name_constraints` -
  (Optional)
  Describes the X.509 name constraints extension.
  Structure is [documented below](#nested_predefined_values_name_constraints).


<a name="nested_predefined_values_key_usage"></a>The `key_usage` block supports:

* `base_key_usage` -
  (Optional)
  Describes high-level ways in which a key may be used.
  Structure is [documented below](#nested_predefined_values_key_usage_base_key_usage).

* `extended_key_usage` -
  (Optional)
  Detailed scenarios in which a key may be used.
  Structure is [documented below](#nested_predefined_values_key_usage_extended_key_usage).

* `unknown_extended_key_usages` -
  (Optional)
  Used to describe extended key usages that are not listed in the KeyUsage.ExtendedKeyUsageOptions message.
  Structure is [documented below](#nested_predefined_values_key_usage_unknown_extended_key_usages).


<a name="nested_predefined_values_key_usage_base_key_usage"></a>The `base_key_usage` block supports:

* `digital_signature` -
  (Optional)
  The key may be used for digital signatures.

* `content_commitment` -
  (Optional)
  The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation".

* `key_encipherment` -
  (Optional)
  The key may be used to encipher other keys.

* `data_encipherment` -
  (Optional)
  The key may be used to encipher data.

* `key_agreement` -
  (Optional)
  The key may be used in a key agreement protocol.

* `cert_sign` -
  (Optional)
  The key may be used to sign certificates.

* `crl_sign` -
  (Optional)
  The key may be used sign certificate revocation lists.

* `encipher_only` -
  (Optional)
  The key may be used to encipher only.

* `decipher_only` -
  (Optional)
  The key may be used to decipher only.

<a name="nested_predefined_values_key_usage_extended_key_usage"></a>The `extended_key_usage` block supports:

* `server_auth` -
  (Optional)
  Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS.

* `client_auth` -
  (Optional)
  Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS.

* `code_signing` -
  (Optional)
  Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication".

* `email_protection` -
  (Optional)
  Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection".

* `time_stamping` -
  (Optional)
  Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time".

* `ocsp_signing` -
  (Optional)
  Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses".

<a name="nested_predefined_values_key_usage_unknown_extended_key_usages"></a>The `unknown_extended_key_usages` block supports:

* `object_id_path` -
  (Required)
  Required. The parts of an OID path. The most significant parts of the path come first.

<a name="nested_predefined_values_ca_options"></a>The `ca_options` block supports:

* `is_ca` -
  (Optional)
  Optional. Refers to the "CA" X.509 extension, which is a boolean value. When this value is true, the "CA" in Basic Constraints extension will be set to true.

* `null_ca` -
  (Optional)
  Optional. When true, the "CA" in Basic Constraints extension will be set to null and omitted from the CA certificate.
  If both `is_ca` and `null_ca` are unset, the "CA" in Basic Constraints extension will be set to false.
  Note that the behavior when `is_ca = false` for this resource is different from the behavior in the Certificate Authority, Certificate and CaPool resources.

* `max_issuer_path_length` -
  (Optional)
  Optional. Refers to the "path length constraint" in Basic Constraints extension. For a CA certificate, this value describes the depth of
  subordinate CA certificates that are allowed. If this value is less than 0, the request will fail.

* `zero_max_issuer_path_length` -
  (Optional)
  Optional. When true, the "path length constraint" in Basic Constraints extension will be set to 0.
  if both `max_issuer_path_length` and `zero_max_issuer_path_length` are unset,
  the max path length will be omitted from the CA certificate.

<a name="nested_predefined_values_policy_ids"></a>The `policy_ids` block supports:

* `object_id_path` -
  (Required)
  Required. The parts of an OID path. The most significant parts of the path come first.

<a name="nested_predefined_values_additional_extensions"></a>The `additional_extensions` block supports:

* `object_id` -
  (Required)
  Required. The OID for this X.509 extension.
  Structure is [documented below](#nested_predefined_values_additional_extensions_additional_extensions_object_id).

* `critical` -
  (Optional)
  Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error).

* `value` -
  (Required)
  Required. The value of this X.509 extension.


<a name="nested_predefined_values_additional_extensions_additional_extensions_object_id"></a>The `object_id` block supports:

* `object_id_path` -
  (Required)
  Required. The parts of an OID path. The most significant parts of the path come first.

<a name="nested_predefined_values_name_constraints"></a>The `name_constraints` block supports:

* `critical` -
  (Required)
  Indicates whether or not the name constraints are marked critical.

* `permitted_dns_names` -
  (Optional)
  Contains permitted DNS names. Any DNS name that can be
  constructed by simply adding zero or more labels to
  the left-hand side of the name satisfies the name constraint.
  For example, `example.com`, `www.example.com`, `www.sub.example.com`
  would satisfy `example.com` while `example1.com` does not.

* `excluded_dns_names` -
  (Optional)
  Contains excluded DNS names. Any DNS name that can be
  constructed by simply adding zero or more labels to
  the left-hand side of the name satisfies the name constraint.
  For example, `example.com`, `www.example.com`, `www.sub.example.com`
  would satisfy `example.com` while `example1.com` does not.

* `permitted_ip_ranges` -
  (Optional)
  Contains the permitted IP ranges. For IPv4 addresses, the ranges
  are expressed using CIDR notation as specified in RFC 4632.
  For IPv6 addresses, the ranges are expressed in similar encoding as IPv4
  addresses.

* `excluded_ip_ranges` -
  (Optional)
  Contains the excluded IP ranges. For IPv4 addresses, the ranges
  are expressed using CIDR notation as specified in RFC 4632.
  For IPv6 addresses, the ranges are expressed in similar encoding as IPv4
  addresses.

* `permitted_email_addresses` -
  (Optional)
  Contains the permitted email addresses. The value can be a particular
  email address, a hostname to indicate all email addresses on that host or
  a domain with a leading period (e.g. `.example.com`) to indicate
  all email addresses in that domain.

* `excluded_email_addresses` -
  (Optional)
  Contains the excluded email addresses. The value can be a particular
  email address, a hostname to indicate all email addresses on that host or
  a domain with a leading period (e.g. `.example.com`) to indicate
  all email addresses in that domain.

* `permitted_uris` -
  (Optional)
  Contains the permitted URIs that apply to the host part of the name.
  The value can be a hostname or a domain with a
  leading period (like `.example.com`)

* `excluded_uris` -
  (Optional)
  Contains the excluded URIs that apply to the host part of the name.
  The value can be a hostname or a domain with a
  leading period (like `.example.com`)

<a name="nested_identity_constraints"></a>The `identity_constraints` block supports:

* `cel_expression` -
  (Optional)
  Optional. A CEL expression that may be used to validate the resolved X.509 Subject and/or Subject Alternative Name before a certificate is signed. To see the full allowed syntax and some examples, see https://cloud.google.com/certificate-authority-service/docs/using-cel
  Structure is [documented below](#nested_identity_constraints_cel_expression).

* `allow_subject_passthrough` -
  (Required)
  Required. If this is true, the Subject field may be copied from a certificate request into the signed certificate. Otherwise, the requested Subject will be discarded.

* `allow_subject_alt_names_passthrough` -
  (Required)
  Required. If this is true, the SubjectAltNames extension may be copied from a certificate request into the signed certificate. Otherwise, the requested SubjectAltNames will be discarded.


<a name="nested_identity_constraints_cel_expression"></a>The `cel_expression` block supports:

* `expression` -
  (Optional)
  Textual representation of an expression in Common Expression Language syntax.

* `title` -
  (Optional)
  Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.

* `description` -
  (Optional)
  Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.

* `location` -
  (Optional)
  Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.

<a name="nested_passthrough_extensions"></a>The `passthrough_extensions` block supports:

* `known_extensions` -
  (Optional)
  Optional. A set of named X.509 extensions. Will be combined with additional_extensions to determine the full set of X.509 extensions.

* `additional_extensions` -
  (Optional)
  Optional. A set of ObjectIds identifying custom X.509 extensions. Will be combined with known_extensions to determine the full set of X.509 extensions.
  Structure is [documented below](#nested_passthrough_extensions_additional_extensions).


<a name="nested_passthrough_extensions_additional_extensions"></a>The `additional_extensions` block supports:

* `object_id_path` -
  (Required)
  Required. The parts of an OID path. The most significant parts of the path come first.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}`

* `create_time` -
  Output only. The time at which this CertificateTemplate was created.

* `update_time` -
  Output only. The time at which this CertificateTemplate was updated.

* `terraform_labels` -
  The combination of labels configured directly on the resource
   and default labels configured on the provider.

* `effective_labels` -
  All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


CertificateTemplate can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}`
* `{{project}}/{{location}}/{{name}}`
* `{{location}}/{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import CertificateTemplate using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}"
  to = google_privateca_certificate_template.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), CertificateTemplate can be imported using one of the formats above. For example:

```
$ terraform import google_privateca_certificate_template.default projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}
$ terraform import google_privateca_certificate_template.default {{project}}/{{location}}/{{name}}
$ terraform import google_privateca_certificate_template.default {{location}}/{{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
