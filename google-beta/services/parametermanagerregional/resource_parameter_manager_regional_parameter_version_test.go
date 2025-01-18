// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package parametermanagerregional_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccParameterManagerRegionalRegionalParameterVersion_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckParameterManagerRegionalRegionalParameterVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterManagerRegionalRegionalParameterVersion_basic(context),
			},
			{
				ResourceName:            "google_parameter_manager_regional_parameter_version.regional-parameter-version-update",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parameter", "parameter_version_id"},
			},
			{
				Config: testAccParameterManagerRegionalRegionalParameterVersion_update(context),
			},
			{
				ResourceName:            "google_parameter_manager_regional_parameter_version.regional-parameter-version-update",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parameter", "parameter_version_id"},
			},
			{
				Config: testAccParameterManagerRegionalRegionalParameterVersion_basic(context),
			},
			{
				ResourceName:            "google_parameter_manager_regional_parameter_version.regional-parameter-version-update",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parameter", "parameter_version_id"},
			},
		},
	})
}

func testAccParameterManagerRegionalRegionalParameterVersion_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_parameter_manager_regional_parameter" "regional-parameter-update" {
  provider = google-beta
  parameter_id = "tf_test_regional_parameter%{random_suffix}"
  location = "us-central1"
}

resource "google_parameter_manager_regional_parameter_version" "regional-parameter-version-update" {
  provider = google-beta
  parameter = google_parameter_manager_regional_parameter.regional-parameter-update.id
  parameter_version_id = "tf_test_regional_parameter_version%{random_suffix}"
  parameter_data = "regional-parameter-version-data"
}
`, context)
}

func testAccParameterManagerRegionalRegionalParameterVersion_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_parameter_manager_regional_parameter" "regional-parameter-update" {
  provider = google-beta
  parameter_id = "tf_test_regional_parameter%{random_suffix}"
  location = "us-central1"
}

resource "google_parameter_manager_regional_parameter_version" "regional-parameter-version-update" {
  provider = google-beta
  parameter = google_parameter_manager_regional_parameter.regional-parameter-update.id
  parameter_version_id = "tf_test_regional_parameter_version%{random_suffix}"
  parameter_data = "regional-parameter-version-data"
  disabled = true
}
`, context)
}