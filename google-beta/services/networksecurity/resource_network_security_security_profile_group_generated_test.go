// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package networksecurity_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccNetworkSecuritySecurityProfileGroup_networkSecuritySecurityProfileGroupBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecuritySecurityProfileGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecuritySecurityProfileGroup_networkSecuritySecurityProfileGroupBasicExample(context),
			},
			{
				ResourceName:            "google_network_security_security_profile_group.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "parent", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecuritySecurityProfileGroup_networkSecuritySecurityProfileGroupBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_security_profile_group" "default" {
  name                      = "tf-test-sec-profile-group%{random_suffix}"
  parent                    = "organizations/%{org_id}"
  description               = "my description"
  threat_prevention_profile = google_network_security_security_profile.security_profile.id

  labels = {
    foo = "bar"
  }
}

resource "google_network_security_security_profile" "security_profile" {
    name        = "tf-test-sec-profile%{random_suffix}"
    type        = "THREAT_PREVENTION"
    parent      = "organizations/%{org_id}"
    location    = "global"
}
`, context)
}

func TestAccNetworkSecuritySecurityProfileGroup_networkSecuritySecurityProfileGroupMirroringExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkSecuritySecurityProfileGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecuritySecurityProfileGroup_networkSecuritySecurityProfileGroupMirroringExample(context),
			},
			{
				ResourceName:            "google_network_security_security_profile_group.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "parent", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecuritySecurityProfileGroup_networkSecuritySecurityProfileGroupMirroringExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_network_security_mirroring_deployment_group" "default" {
  provider                      = google-beta
  mirroring_deployment_group_id = "tf-test-deployment-group%{random_suffix}"
  location                      = "global"
  network                       = google_compute_network.default.id
}

resource "google_network_security_mirroring_endpoint_group" "default" {
  provider                      = google-beta
  mirroring_endpoint_group_id   = "tf-test-endpoint-group%{random_suffix}"
  location                      = "global"
  mirroring_deployment_group    = google_network_security_mirroring_deployment_group.default.id
}

resource "google_network_security_security_profile" "default" {
  provider    = google-beta
  name        = "tf-test-sec-profile%{random_suffix}"
  parent      = "organizations/%{org_id}"
  description = "my description"
  type        = "CUSTOM_MIRRORING"

  custom_mirroring_profile {
    mirroring_endpoint_group = google_network_security_mirroring_endpoint_group.default.id
  }
}

resource "google_network_security_security_profile_group" "default" {
  provider                 = google-beta
  name                     = "tf-test-sec-profile-group%{random_suffix}"
  parent                   = "organizations/%{org_id}"
  description              = "my description"
  custom_mirroring_profile = google_network_security_security_profile.default.id
}
`, context)
}

func TestAccNetworkSecuritySecurityProfileGroup_networkSecuritySecurityProfileGroupInterceptExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkSecuritySecurityProfileGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecuritySecurityProfileGroup_networkSecuritySecurityProfileGroupInterceptExample(context),
			},
			{
				ResourceName:            "google_network_security_security_profile_group.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "parent", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecuritySecurityProfileGroup_networkSecuritySecurityProfileGroupInterceptExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_network_security_intercept_deployment_group" "default" {
  provider                      = google-beta
  intercept_deployment_group_id = "tf-test-deployment-group%{random_suffix}"
  location                      = "global"
  network                       = google_compute_network.default.id
}

resource "google_network_security_intercept_endpoint_group" "default" {
  provider                      = google-beta
  intercept_endpoint_group_id   = "tf-test-endpoint-group%{random_suffix}"
  location                      = "global"
  intercept_deployment_group    = google_network_security_intercept_deployment_group.default.id
}

resource "google_network_security_security_profile" "default" {
  provider    = google-beta
  name        = "tf-test-sec-profile%{random_suffix}"
  parent      = "organizations/%{org_id}"
  description = "my description"
  type        = "CUSTOM_INTERCEPT"

  custom_intercept_profile {
    intercept_endpoint_group = google_network_security_intercept_endpoint_group.default.id
  }
}

resource "google_network_security_security_profile_group" "default" {
  provider                 = google-beta
  name                     = "tf-test-sec-profile-group%{random_suffix}"
  parent                   = "organizations/%{org_id}"
  description              = "my description"
  custom_intercept_profile = google_network_security_security_profile.default.id
}
`, context)
}

func testAccCheckNetworkSecuritySecurityProfileGroupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_security_security_profile_group" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/securityProfileGroups/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("NetworkSecuritySecurityProfileGroup still exists at %s", url)
			}
		}

		return nil
	}
}
