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

package securesourcemanager_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccSecureSourceManagerBranchRule_secureSourceManagerBranchRuleBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"prevent_destroy": false,
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecureSourceManagerBranchRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecureSourceManagerBranchRule_secureSourceManagerBranchRuleBasicExample(context),
			},
			{
				ResourceName:            "google_secure_source_manager_branch_rule.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"branch_rule_id", "location", "repository_id"},
			},
		},
	})
}

func testAccSecureSourceManagerBranchRule_secureSourceManagerBranchRuleBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secure_source_manager_instance" "instance" {
    location = "us-central1"
    instance_id = "tf-test-my-basic-instance%{random_suffix}"
    # Prevent accidental deletions.
    lifecycle {
        prevent_destroy = "%{prevent_destroy}"
    }
}

resource "google_secure_source_manager_repository" "repository" {
    repository_id = "tf-test-my-basic-repository%{random_suffix}"
    location = google_secure_source_manager_instance.instance.location
    instance = google_secure_source_manager_instance.instance.name
    # Prevent accidental deletions.
    lifecycle {
        prevent_destroy = "%{prevent_destroy}"
    }
}

resource "google_secure_source_manager_branch_rule" "basic" {
    branch_rule_id = "tf-test-my-basic-branchrule%{random_suffix}"
    repository_id = google_secure_source_manager_repository.repository.repository_id
    location = google_secure_source_manager_repository.repository.location
    # This field is required for BranchRule creation
    include_pattern = "main"
}
`, context)
}

func TestAccSecureSourceManagerBranchRule_secureSourceManagerBranchRuleWithFieldsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"prevent_destroy": false,
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecureSourceManagerBranchRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecureSourceManagerBranchRule_secureSourceManagerBranchRuleWithFieldsExample(context),
			},
			{
				ResourceName:            "google_secure_source_manager_branch_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"branch_rule_id", "location", "repository_id"},
			},
		},
	})
}

func testAccSecureSourceManagerBranchRule_secureSourceManagerBranchRuleWithFieldsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secure_source_manager_instance" "instance" {
    location = "us-central1"
    instance_id = "tf-test-my-initial-instance%{random_suffix}"
    # Prevent accidental deletions.
    lifecycle {
        prevent_destroy = "%{prevent_destroy}"
    }
}

resource "google_secure_source_manager_repository" "repository" {
    repository_id = "tf-test-my-initial-repository%{random_suffix}"
    instance = google_secure_source_manager_instance.instance.name
    location = google_secure_source_manager_instance.instance.location
    # Prevent accidental deletions.
    lifecycle {
        prevent_destroy = "%{prevent_destroy}"
    }
}

resource "google_secure_source_manager_branch_rule" "default" {
    branch_rule_id = "tf-test-my-initial-branchrule%{random_suffix}"
    location = google_secure_source_manager_repository.repository.location
    repository_id = google_secure_source_manager_repository.repository.repository_id
    include_pattern = "test"
    minimum_approvals_count   = 2
    minimum_reviews_count     = 2
    require_comments_resolved = true
    require_linear_history    = true
    require_pull_request      = true
    disabled = false
    allow_stale_reviews = false
}
`, context)
}

func testAccCheckSecureSourceManagerBranchRuleDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_secure_source_manager_branch_rule" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{SecureSourceManagerBasePath}}projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}/branchRules/{{branch_rule_id}}")
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
				return fmt.Errorf("SecureSourceManagerBranchRule still exists at %s", url)
			}
		}

		return nil
	}
}
