// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package chronicle_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccChronicleDataAccessLabel_chronicleDataaccesslabelBasicExample_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"chronicle_id":  envvar.GetTestChronicleInstanceIdFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckChronicleDataAccessLabelDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccChronicleDataAccessLabel_chronicleDataaccesslabelBasicExample_full(context),
			},
			{
				ResourceName:            "google_chronicle_data_access_label.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"data_access_label_id", "instance", "location"},
			},

			{
				Config: testAccChronicleDataAccessLabel_chronicleDataaccesslabelBasicExample_update(context),
			},
			{
				ResourceName:            "google_chronicle_data_access_label.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"data_access_label_id", "instance", "location"},
			},
		},
	})
}

func testAccChronicleDataAccessLabel_chronicleDataaccesslabelBasicExample_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_chronicle_data_access_label" "example" {
  provider = "google-beta"
  location = "us" 
  instance = "%{chronicle_id}"
  data_access_label_id = "tf-test-label-id%{random_suffix}"
  udm_query = "principal.hostname=\"google.com\""
  description = "tf-test-label-description%{random_suffix}"
}
`, context)
}

func testAccChronicleDataAccessLabel_chronicleDataaccesslabelBasicExample_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_chronicle_data_access_label" "example" {
  provider = "google-beta"
  location = "us" 
  instance = "%{chronicle_id}"
  data_access_label_id = "tf-test-label-id%{random_suffix}"
  udm_query = "principal.hostname=\"google-updated.com\""
  description = "tf-test-label-updated-description%{random_suffix}"
}
`, context)
}