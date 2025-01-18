// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package parametermanagerregional_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccDataSourceParameterManagerRegionalRegionalParameters_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckParameterManagerRegionalRegionalParameterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceParameterManagerRegionalRegionalParameters_basic(context),
				Check: resource.ComposeTestCheckFunc(
					checkListDataSourceStateMatchesResourceStateWithIgnores(
						"data.google_parameter_manager_regional_parameters.regional-parameters-datasource",
						"google_parameter_manager_regional_parameter.regional-parameters",
						map[string]struct{}{
							"id":      {},
							"project": {},
						},
					),
				),
			},
		},
	})
}

func testAccDataSourceParameterManagerRegionalRegionalParameters_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google-beta" {
  add_terraform_attribution_label = false
}

resource "google_parameter_manager_regional_parameter" "regional-parameters" {
  provider = google-beta
  parameter_id = "tf_test_regional_parameter%{random_suffix}"
  format = "YAML"
  location = "us-central1"

  labels = {
    key1 = "val1"
    key2 = "val2"
    key3 = "val3"
    key4 = "val4"
    key5 = "val5"
  }
}

data "google_parameter_manager_regional_parameters" "regional-parameters-datasource" {
  provider = google-beta
  depends_on = [
    google_parameter_manager_regional_parameter.regional-parameters
  ]
  location = "us-central1"
}
`, context)
}

func TestAccDataSourceParameterManagerRegionalRegionalParameters_filter(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckParameterManagerRegionalRegionalParameterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceParameterManagerRegionalRegionalParameters_filter(context),
				Check: resource.ComposeTestCheckFunc(
					checkListDataSourceStateMatchesResourceStateWithIgnoresForAppliedFilter(
						"data.google_parameter_manager_regional_parameters.regional-parameters-datasource-filter",
						"google_parameter_manager_regional_parameter.regional-parameters-1",
						"google_parameter_manager_regional_parameter.regional-parameters-2",
						map[string]struct{}{
							"id":      {},
							"project": {},
						},
					),
				),
			},
		},
	})
}

func testAccDataSourceParameterManagerRegionalRegionalParameters_filter(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google-beta" {
  add_terraform_attribution_label = false
}

resource "google_parameter_manager_regional_parameter" "regional-parameters-1" {
  provider = google-beta
  parameter_id = "tf_test_regional_parameter%{random_suffix}"
  format = "JSON"
  location = "us-central1"

  labels = {
    key1 = "val1"
  }
}

resource "google_parameter_manager_regional_parameter" "regional-parameters-2" {
  provider = google-beta
  parameter_id = "tf_test_regional_parameter_2_%{random_suffix}"
  format = "YAML"
  location = "us-central1"

  labels = {
    keyoth1 = "valoth1"
  }
}

data "google_parameter_manager_regional_parameters" "regional-parameters-datasource-filter" {
  provider = google-beta
  filter = "format:JSON"
  location = "us-central1"

  depends_on = [
    google_parameter_manager_regional_parameter.regional-parameters-1,
	google_parameter_manager_regional_parameter.regional-parameters-2
  ]
}
`, context)
}

// This function checks data source state matches for resourceName parameter manager regional parameter state
func checkListDataSourceStateMatchesResourceStateWithIgnores(dataSourceName, resourceName string, ignoreFields map[string]struct{}) func(*terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[dataSourceName]
		if !ok {
			return fmt.Errorf("can't find %s in state", dataSourceName)
		}

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("can't find %s in state", resourceName)
		}

		dsAttr := ds.Primary.Attributes
		rsAttr := rs.Primary.Attributes

		err := checkFieldsMatchForDataSourceStateAndResourceState(dsAttr, rsAttr, ignoreFields)
		if err != nil {
			return err
		}
		return nil
	}
}

// This function checks whether all the attributes of the parameter manager regional parameter resource and the attributes of the parameter manager regional parameter inside the data source list are the same
func checkFieldsMatchForDataSourceStateAndResourceState(dsAttr, rsAttr map[string]string, ignoreFields map[string]struct{}) error {
	totalParameters, err := strconv.Atoi(dsAttr["parameters.#"])
	if err != nil {
		return errors.New("couldn't convert length of regional parameters list to integer")
	}
	index := "-1"
	for i := 0; i < totalParameters; i++ {
		if dsAttr["parameters."+strconv.Itoa(i)+".name"] == rsAttr["name"] {
			index = strconv.Itoa(i)
		}
	}

	if index == "-1" {
		return errors.New("the newly created regional parameter is not found in the data source")
	}

	errMsg := ""
	// Data sources are often derived from resources, so iterate over the resource fields to
	// make sure all fields are accounted for in the data source.
	// If a field exists in the data source but not in the resource, its expected value should
	// be checked separately.
	for k := range rsAttr {
		if _, ok := ignoreFields[k]; ok {
			continue
		}
		if k == "%" {
			continue
		}
		if dsAttr["parameters."+index+"."+k] != rsAttr[k] {
			// ignore data sources where an empty list is being compared against a null list.
			if k[len(k)-1:] == "#" && (dsAttr["parameters."+index+"."+k] == "" || dsAttr["parameters."+index+"."+k] == "0") && (rsAttr[k] == "" || rsAttr[k] == "0") {
				continue
			}
			errMsg += fmt.Sprintf("%s is %s; want %s\n", k, dsAttr["parameters."+index+"."+k], rsAttr[k])
		}
	}

	if errMsg != "" {
		return errors.New(errMsg)
	}

	return nil
}

// This function checks state match for resourceName and asserts the absense of resourceName2 in data source
func checkListDataSourceStateMatchesResourceStateWithIgnoresForAppliedFilter(dataSourceName, resourceName, resourceName2 string, ignoreFields map[string]struct{}) func(*terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[dataSourceName]
		if !ok {
			return fmt.Errorf("can't find %s in state", dataSourceName)
		}

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("can't find %s in state", resourceName)
		}

		rs2, ok := s.RootModule().Resources[resourceName2]
		if !ok {
			return fmt.Errorf("can't find %s in state", resourceName2)
		}

		dsAttr := ds.Primary.Attributes
		rsAttr := rs.Primary.Attributes
		rsAttr2 := rs2.Primary.Attributes

		err := checkFieldsMatchForDataSourceStateAndResourceState(dsAttr, rsAttr, ignoreFields)
		if err != nil {
			return err
		}
		err = checkResourceAbsentInDataSourceAfterFilterApplied(dsAttr, rsAttr2)
		return err
	}
}

// This function asserts the absence of the parameter manager regional parameter resource which would not be included in the data source list due to the filter applied.
func checkResourceAbsentInDataSourceAfterFilterApplied(dsAttr, rsAttr map[string]string) error {
	totalParameters, err := strconv.Atoi(dsAttr["parameters.#"])
	if err != nil {
		return errors.New("couldn't convert length of regional parameters list to integer")
	}
	for i := 0; i < totalParameters; i++ {
		if dsAttr["parameters."+strconv.Itoa(i)+".name"] == rsAttr["name"] {
			return errors.New("the resource is present in the data source even after the filter is applied")
		}
	}
	return nil
}