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

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccNetworkSecurityAddressGroup_networkSecurityAddressGroupsBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityAddressGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityAddressGroup_networkSecurityAddressGroupsBasicExample(context),
			},
			{
				ResourceName:            "google_network_security_address_group.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccNetworkSecurityAddressGroup_networkSecurityAddressGroupsBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_network_security_address_group" "default" {
  provider    = google-beta
  name        = "tf-test-my-address-groups%{random_suffix}"
  location    = "us-central1"
  type        = "IPV4"
  capacity    = "100"
  items       = ["208.80.154.224/32"]
}
`, context)
}

func TestAccNetworkSecurityAddressGroup_networkSecurityAddressGroupsAdvancedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityAddressGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityAddressGroup_networkSecurityAddressGroupsAdvancedExample(context),
			},
			{
				ResourceName:            "google_network_security_address_group.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccNetworkSecurityAddressGroup_networkSecurityAddressGroupsAdvancedExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_network_security_address_group" "default" {
  provider    = google-beta
  name        = "tf-test-my-address-groups%{random_suffix}"
  location    = "us-central1"
  description = "my description"
  type        = "IPV4"
  capacity    = "100"
  items       = ["208.80.154.224/32"]
}
`, context)
}

func testAccCheckNetworkSecurityAddressGroupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_security_address_group" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/addressGroups/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("NetworkSecurityAddressGroup still exists at %s", url)
			}
		}

		return nil
	}
}