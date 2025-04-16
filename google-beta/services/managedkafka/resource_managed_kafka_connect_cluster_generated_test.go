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

package managedkafka_test

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

func TestAccManagedKafkaConnectCluster_managedkafkaConnectClusterBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccCheckManagedKafkaConnectClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccManagedKafkaConnectCluster_managedkafkaConnectClusterBasicExample(context),
			},
			{
				ResourceName:            "google_managed_kafka_connect_cluster.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connect_cluster_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccManagedKafkaConnectCluster_managedkafkaConnectClusterBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "project" {
  project_id      = "tf-test%{random_suffix}"
  name            = "tf-test%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
  deletion_policy = "DELETE"

  provider = google-beta
}

resource "time_sleep" "wait_60_seconds" {
  create_duration = "60s"
  depends_on = [google_project.project]
}

resource "google_project_service" "compute" {
  project = google_project.project.project_id
  service = "compute.googleapis.com"
  depends_on = [time_sleep.wait_60_seconds]

  provider = google-beta
}

resource "google_project_service" "managedkafka" {
  project = google_project.project.project_id
  service = "managedkafka.googleapis.com"
  depends_on = [google_project_service.compute]

  provider = google-beta
}

resource "time_sleep" "wait_120_seconds" {
  create_duration = "120s"
  depends_on = [google_project_service.managedkafka]
}

resource "google_compute_subnetwork" "mkc_secondary_subnet" {
  project       = google_project.project.project_id
  name          = "tf-test-my-secondary-subnetwork%{random_suffix}"
  ip_cidr_range = "10.3.0.0/16"
  region        = "us-central1"
  network       = "default"
  depends_on = [time_sleep.wait_120_seconds]

  provider = google-beta
}

resource "google_managed_kafka_cluster" "gmk_cluster" {
  project = google_project.project.project_id
  cluster_id = "tf-test-my-cluster%{random_suffix}"
  location = "us-central1"
  capacity_config {
    vcpu_count = 3
    memory_bytes = 3221225472
  }
  gcp_config {
    access_config {
      network_configs {
        subnet = "projects/${google_project.project.project_id}/regions/us-central1/subnetworks/default"
      }
    }
  }
  depends_on = [google_project_service.managedkafka]

  provider = google-beta
}

resource "google_managed_kafka_connect_cluster" "example" {
  project = google_project.project.project_id
  connect_cluster_id = "tf-test-my-connect-cluster%{random_suffix}"
  kafka_cluster = "projects/${google_project.project.project_id}/locations/us-central1/clusters/${google_managed_kafka_cluster.gmk_cluster.cluster_id}"
  location = "us-central1"
  capacity_config {
    vcpu_count = 12
    memory_bytes = 21474836480
  }
  gcp_config {
    access_config {
      network_configs {
        primary_subnet = "projects/${google_project.project.project_id}/regions/us-central1/subnetworks/default"
        additional_subnets = ["${google_compute_subnetwork.mkc_secondary_subnet.id}"]
        dns_domain_names = ["${google_managed_kafka_cluster.gmk_cluster.cluster_id}.us-central1.managedkafka.${google_project.project.project_id}.cloud.goog"]
      }
    }
  }
  labels = {
    key = "value"
  }
  depends_on = [google_project_service.managedkafka]

  provider = google-beta
}
`, context)
}

func testAccCheckManagedKafkaConnectClusterDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_managed_kafka_connect_cluster" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ManagedKafkaBasePath}}projects/{{project}}/locations/{{location}}/connectClusters/{{connect_cluster_id}}")
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
				return fmt.Errorf("ManagedKafkaConnectCluster still exists at %s", url)
			}
		}

		return nil
	}
}
