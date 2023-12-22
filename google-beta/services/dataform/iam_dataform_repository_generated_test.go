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

package dataform_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccDataformRepositoryIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataformRepositoryIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataform_repository_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/repositories/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf_test_dataform_repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccDataformRepositoryIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_dataform_repository_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/repositories/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf_test_dataform_repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataformRepositoryIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccDataformRepositoryIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataform_repository_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/repositories/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf_test_dataform_repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataformRepositoryIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataformRepositoryIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_dataform_repository_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_dataform_repository_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/repositories/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf_test_dataform_repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataformRepositoryIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_dataform_repository_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/repositories/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf_test_dataform_repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataformRepositoryIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sourcerepo_repository" "git_repository" {
  provider = google-beta
  name = "my/repository%{random_suffix}"
}

resource "google_secret_manager_secret" "secret" {
  provider = google-beta
  secret_id = "tf-test-my-secret%{random_suffix}"

  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "secret_version" {
  provider = google-beta
  secret = google_secret_manager_secret.secret.id

  secret_data = "tf-test-secret-data%{random_suffix}"
}

resource "google_dataform_repository" "dataform_respository" {
  provider = google-beta
  name = "tf_test_dataform_repository%{random_suffix}"
  display_name = "tf_test_dataform_repository%{random_suffix}"
  npmrc_environment_variables_secret_version = google_secret_manager_secret_version.secret_version.id

  labels = {
    label_foo1 = "label-bar1"
  }

  git_remote_settings {
      url = google_sourcerepo_repository.git_repository.url
      default_branch = "main"
      authentication_token_secret_version = google_secret_manager_secret_version.secret_version.id
  }

  workspace_compilation_overrides {
    default_database = "database"
    schema_suffix = "_suffix"
    table_prefix = "prefix_"
  }
}

resource "google_dataform_repository_iam_member" "foo" {
  provider = google-beta
  project = google_dataform_repository.dataform_respository.project
  region = google_dataform_repository.dataform_respository.region
  repository = google_dataform_repository.dataform_respository.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccDataformRepositoryIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sourcerepo_repository" "git_repository" {
  provider = google-beta
  name = "my/repository%{random_suffix}"
}

resource "google_secret_manager_secret" "secret" {
  provider = google-beta
  secret_id = "tf-test-my-secret%{random_suffix}"

  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "secret_version" {
  provider = google-beta
  secret = google_secret_manager_secret.secret.id

  secret_data = "tf-test-secret-data%{random_suffix}"
}

resource "google_dataform_repository" "dataform_respository" {
  provider = google-beta
  name = "tf_test_dataform_repository%{random_suffix}"
  display_name = "tf_test_dataform_repository%{random_suffix}"
  npmrc_environment_variables_secret_version = google_secret_manager_secret_version.secret_version.id

  labels = {
    label_foo1 = "label-bar1"
  }

  git_remote_settings {
      url = google_sourcerepo_repository.git_repository.url
      default_branch = "main"
      authentication_token_secret_version = google_secret_manager_secret_version.secret_version.id
  }

  workspace_compilation_overrides {
    default_database = "database"
    schema_suffix = "_suffix"
    table_prefix = "prefix_"
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_dataform_repository_iam_policy" "foo" {
  provider = google-beta
  project = google_dataform_repository.dataform_respository.project
  region = google_dataform_repository.dataform_respository.region
  repository = google_dataform_repository.dataform_respository.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_dataform_repository_iam_policy" "foo" {
  provider = google-beta
  project = google_dataform_repository.dataform_respository.project
  region = google_dataform_repository.dataform_respository.region
  repository = google_dataform_repository.dataform_respository.name
  depends_on = [
    google_dataform_repository_iam_policy.foo
  ]
}
`, context)
}

func testAccDataformRepositoryIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sourcerepo_repository" "git_repository" {
  provider = google-beta
  name = "my/repository%{random_suffix}"
}

resource "google_secret_manager_secret" "secret" {
  provider = google-beta
  secret_id = "tf-test-my-secret%{random_suffix}"

  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "secret_version" {
  provider = google-beta
  secret = google_secret_manager_secret.secret.id

  secret_data = "tf-test-secret-data%{random_suffix}"
}

resource "google_dataform_repository" "dataform_respository" {
  provider = google-beta
  name = "tf_test_dataform_repository%{random_suffix}"
  display_name = "tf_test_dataform_repository%{random_suffix}"
  npmrc_environment_variables_secret_version = google_secret_manager_secret_version.secret_version.id

  labels = {
    label_foo1 = "label-bar1"
  }

  git_remote_settings {
      url = google_sourcerepo_repository.git_repository.url
      default_branch = "main"
      authentication_token_secret_version = google_secret_manager_secret_version.secret_version.id
  }

  workspace_compilation_overrides {
    default_database = "database"
    schema_suffix = "_suffix"
    table_prefix = "prefix_"
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_dataform_repository_iam_policy" "foo" {
  provider = google-beta
  project = google_dataform_repository.dataform_respository.project
  region = google_dataform_repository.dataform_respository.region
  repository = google_dataform_repository.dataform_respository.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataformRepositoryIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sourcerepo_repository" "git_repository" {
  provider = google-beta
  name = "my/repository%{random_suffix}"
}

resource "google_secret_manager_secret" "secret" {
  provider = google-beta
  secret_id = "tf-test-my-secret%{random_suffix}"

  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "secret_version" {
  provider = google-beta
  secret = google_secret_manager_secret.secret.id

  secret_data = "tf-test-secret-data%{random_suffix}"
}

resource "google_dataform_repository" "dataform_respository" {
  provider = google-beta
  name = "tf_test_dataform_repository%{random_suffix}"
  display_name = "tf_test_dataform_repository%{random_suffix}"
  npmrc_environment_variables_secret_version = google_secret_manager_secret_version.secret_version.id

  labels = {
    label_foo1 = "label-bar1"
  }

  git_remote_settings {
      url = google_sourcerepo_repository.git_repository.url
      default_branch = "main"
      authentication_token_secret_version = google_secret_manager_secret_version.secret_version.id
  }

  workspace_compilation_overrides {
    default_database = "database"
    schema_suffix = "_suffix"
    table_prefix = "prefix_"
  }
}

resource "google_dataform_repository_iam_binding" "foo" {
  provider = google-beta
  project = google_dataform_repository.dataform_respository.project
  region = google_dataform_repository.dataform_respository.region
  repository = google_dataform_repository.dataform_respository.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccDataformRepositoryIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sourcerepo_repository" "git_repository" {
  provider = google-beta
  name = "my/repository%{random_suffix}"
}

resource "google_secret_manager_secret" "secret" {
  provider = google-beta
  secret_id = "tf-test-my-secret%{random_suffix}"

  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "secret_version" {
  provider = google-beta
  secret = google_secret_manager_secret.secret.id

  secret_data = "tf-test-secret-data%{random_suffix}"
}

resource "google_dataform_repository" "dataform_respository" {
  provider = google-beta
  name = "tf_test_dataform_repository%{random_suffix}"
  display_name = "tf_test_dataform_repository%{random_suffix}"
  npmrc_environment_variables_secret_version = google_secret_manager_secret_version.secret_version.id

  labels = {
    label_foo1 = "label-bar1"
  }

  git_remote_settings {
      url = google_sourcerepo_repository.git_repository.url
      default_branch = "main"
      authentication_token_secret_version = google_secret_manager_secret_version.secret_version.id
  }

  workspace_compilation_overrides {
    default_database = "database"
    schema_suffix = "_suffix"
    table_prefix = "prefix_"
  }
}

resource "google_dataform_repository_iam_binding" "foo" {
  provider = google-beta
  project = google_dataform_repository.dataform_respository.project
  region = google_dataform_repository.dataform_respository.region
  repository = google_dataform_repository.dataform_respository.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}