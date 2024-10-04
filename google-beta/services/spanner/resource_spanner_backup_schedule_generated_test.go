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

package spanner_test

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

func TestAccSpannerBackupSchedule_spannerBackupScheduleDailyFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSpannerBackupScheduleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSpannerBackupSchedule_spannerBackupScheduleDailyFullExample(context),
			},
			{
				ResourceName:            "google_spanner_backup_schedule.full-backup",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"database", "instance"},
			},
		},
	})
}

func testAccSpannerBackupSchedule_spannerBackupScheduleDailyFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_spanner_instance" "main" {
  name         = "tf-test-instance-id%{random_suffix}"
  config       = "regional-europe-west1"
  display_name = "main-instance"
  num_nodes    = 1
}

resource "google_spanner_database" "database" {
  instance = google_spanner_instance.main.name
  name     = "tf-test-database-id%{random_suffix}"
  version_retention_period = "3d"
  ddl = [
    "CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)",
    "CREATE TABLE t2 (t2 INT64 NOT NULL,) PRIMARY KEY(t2)",
  ]
  deletion_protection = "%{deletion_protection}"
}

resource "google_spanner_backup_schedule" "full-backup" {
  instance = google_spanner_instance.main.name

  database = google_spanner_database.database.name

  name = "tf-test-backup-schedule-id%{random_suffix}"

  retention_duration = "31620000s" // 366 days (maximum possible retention)

  spec {
    cron_spec {
      //   0 2/12 * * * : every 12 hours at (2, 14) hours past midnight in UTC.
      //   0 2,14 * * * : every 12 hours at (2,14) hours past midnight in UTC.
      //   0 2 * * *    : once a day at 2 past midnight in UTC.
      //   0 2 * * 0    : once a week every Sunday at 2 past midnight in UTC.
      //   0 2 8 * *    : once a month on 8th day at 2 past midnight in UTC.
      text = "0 12 * * *"
    }
  }
  // The schedule creates only full backups.
  full_backup_spec {}
}
`, context)
}

func TestAccSpannerBackupSchedule_spannerBackupScheduleDailyIncrementalExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSpannerBackupScheduleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSpannerBackupSchedule_spannerBackupScheduleDailyIncrementalExample(context),
			},
			{
				ResourceName:            "google_spanner_backup_schedule.incremental-backup",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"database", "instance"},
			},
		},
	})
}

func testAccSpannerBackupSchedule_spannerBackupScheduleDailyIncrementalExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_spanner_instance" "main" {
  name         = "tf-test-instance-id%{random_suffix}"
  config       = "regional-europe-west1"
  display_name = "main-instance"
  num_nodes    = 1
  edition      = "ENTERPRISE"
}

resource "google_spanner_database" "database" {
  instance = google_spanner_instance.main.name
  name     = "tf-test-database-id%{random_suffix}"
  version_retention_period = "3d"
  ddl = [
    "CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)",
    "CREATE TABLE t2 (t2 INT64 NOT NULL,) PRIMARY KEY(t2)",
  ]
  deletion_protection = "%{deletion_protection}"
}

resource "google_spanner_backup_schedule" "incremental-backup" {
  instance = google_spanner_instance.main.name

  database = google_spanner_database.database.name

  name = "tf-test-backup-schedule-id%{random_suffix}"
  
  retention_duration = "31620000s" // 366 days (maximum possible retention)

  spec {
    cron_spec {
      //   0 2/12 * * * : every 12 hours at (2, 14) hours past midnight in UTC.
      //   0 2,14 * * * : every 12 hours at (2,14) hours past midnight in UTC.
      //   0 2 * * *    : once a day at 2 past midnight in UTC.
      //   0 2 * * 0    : once a week every Sunday at 2 past midnight in UTC.
      //   0 2 8 * *    : once a month on 8th day at 2 past midnight in UTC.
      text = "0 12 * * *"
    }
  }
  // The schedule creates incremental backup chains.
  incremental_backup_spec {}
}
`, context)
}

func testAccCheckSpannerBackupScheduleDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_spanner_backup_schedule" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{SpannerBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{database}}/backupSchedules/{{name}}")
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
				return fmt.Errorf("SpannerBackupSchedule still exists at %s", url)
			}
		}

		return nil
	}
}
