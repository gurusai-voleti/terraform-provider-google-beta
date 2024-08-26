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

package cloudtasks_test

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

func TestAccCloudTasksQueue_queueBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudTasksQueueDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudTasksQueue_queueBasicExample(context),
			},
			{
				ResourceName:            "google_cloud_tasks_queue.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccCloudTasksQueue_queueBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_tasks_queue" "default" {
  name = "tf-test-cloud-tasks-queue-test%{random_suffix}"
  location = "us-central1"
}
`, context)
}

func TestAccCloudTasksQueue_cloudTasksQueueAdvancedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudTasksQueueDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudTasksQueue_cloudTasksQueueAdvancedExample(context),
			},
			{
				ResourceName:            "google_cloud_tasks_queue.advanced_configuration",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"app_engine_routing_override.0.instance", "app_engine_routing_override.0.service", "app_engine_routing_override.0.version", "location"},
			},
		},
	})
}

func testAccCloudTasksQueue_cloudTasksQueueAdvancedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_tasks_queue" "advanced_configuration" {
  name = "tf-test-instance-name%{random_suffix}"
  location = "us-central1"

  app_engine_routing_override {
    service = "worker"
    version = "1.0"
    instance = "test"
  }

  rate_limits {
    max_concurrent_dispatches = 3
    max_dispatches_per_second = 2
  }

  retry_config {
    max_attempts = 5
    max_retry_duration = "4s"
    max_backoff = "3s"
    min_backoff = "2s"
    max_doublings = 1
  }

  stackdriver_logging_config {
    sampling_ratio = 0.9
  }
}
`, context)
}

func TestAccCloudTasksQueue_cloudTasksQueueHttpTargetOidcExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudTasksQueueDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudTasksQueue_cloudTasksQueueHttpTargetOidcExample(context),
			},
			{
				ResourceName:            "google_cloud_tasks_queue.http_target_oidc",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccCloudTasksQueue_cloudTasksQueueHttpTargetOidcExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_tasks_queue" "http_target_oidc" {
  name     = "tf-test-cloud-tasks-queue-http-target-oidc%{random_suffix}"
  location = "us-central1"

  http_target {
    http_method = "POST"
    uri_override {
      scheme = "HTTPS"
      host   = "oidc.example.com"
      port   = 8443
      path_override {
        path = "/users/1234"
      }
      query_override {
        query_params = "qparam1=123&qparam2=456"
      }
      uri_override_enforce_mode = "IF_NOT_EXISTS"
    }
    header_overrides {
      header {
        key   = "AddSomethingElse"
        value = "MyOtherValue"
      }
    }
    header_overrides {
      header {
        key   = "AddMe"
        value = "MyValue"
      }
    }
    oidc_token {
      service_account_email = google_service_account.oidc_service_account.email
      audience              = "https://oidc.example.com"
    }
  }
}

resource "google_service_account" "oidc_service_account" {
  account_id   = "example-oidc"
  display_name = "Tasks Queue OIDC Service Account"
}
`, context)
}

func TestAccCloudTasksQueue_cloudTasksQueueHttpTargetOauthExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudTasksQueueDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudTasksQueue_cloudTasksQueueHttpTargetOauthExample(context),
			},
			{
				ResourceName:            "google_cloud_tasks_queue.http_target_oauth",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccCloudTasksQueue_cloudTasksQueueHttpTargetOauthExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_tasks_queue" "http_target_oauth" {
  name     = "tf-test-cloud-tasks-queue-http-target-oauth%{random_suffix}"
  location = "us-central1"

  http_target {
    http_method = "POST"
    uri_override {
      scheme = "HTTPS"
      host   = "oauth.example.com"
      port   = 8443
      path_override {
        path = "/users/1234"
      }
      query_override {
        query_params = "qparam1=123&qparam2=456"
      }
      uri_override_enforce_mode = "IF_NOT_EXISTS"
    }
    header_overrides {
      header {
        key   = "AddSomethingElse"
        value = "MyOtherValue"
      }
    }
    header_overrides {
      header {
        key   = "AddMe"
        value = "MyValue"
      }
    }
    oauth_token {
      service_account_email = google_service_account.oauth_service_account.email
      scope                 = "openid https://www.googleapis.com/auth/userinfo.email"
    }
  }
}

resource "google_service_account" "oauth_service_account" {
  account_id   = "example-oauth"
  display_name = "Tasks Queue OAuth Service Account"
}
`, context)
}

func testAccCheckCloudTasksQueueDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloud_tasks_queue" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{CloudTasksBasePath}}projects/{{project}}/locations/{{location}}/queues/{{name}}")
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
				return fmt.Errorf("CloudTasksQueue still exists at %s", url)
			}
		}

		return nil
	}
}
