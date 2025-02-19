package osconfigv2_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccOSConfigV2PolicyOrchestratorForOrganization_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckOSConfigV2PolicyOrchestratorForOrganizationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccOSConfigV2PolicyOrchestratorForOrganization_basic(context),
			},
			{
				ResourceName:            "google_os_config_v2_policy_orchestrator_for_organization.policy_orchestrator_for_organization",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "organization_id", "policy_orchestrator_id", "terraform_labels"},
			},
		},
	})
}

func testAccOSConfigV2PolicyOrchestratorForOrganization_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_os_config_v2_policy_orchestrator_for_organization" "policy_orchestrator_for_organization" {
    provider = google-beta
    policy_orchestrator_id = "tf-test-test-po-org%{random_suffix}"
    organization_id = "%{org_id}"
    
    state = "ACTIVE"
    action = "UPSERT"
    
    orchestrated_resource {
        id = "tf-test-test-orchestrated-resource-org%{random_suffix}"
        os_policy_assignment_v1_payload {
            os_policies {
                id = "tf-test-test-os-policy-org%{random_suffix}"
                mode = "VALIDATION"
                resource_groups {
                    resources {
                        id = "resource-tf"
                        file {
                            content = "file-content-tf"
                            path = "file-path-tf-1"
                            state = "PRESENT"
                        }
                    }
                }
            }
            instance_filter {
                inventories {
                    os_short_name = "windows-10"
                }
            }
            rollout {
                disruption_budget {
                    percent = 100
                }
                min_wait_duration = "60s"
            }
        }
    }
}
`, context)
}

func testAccOSConfigV2PolicyOrchestratorForOrganization_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_os_config_v2_policy_orchestrator_for_organization" "policy_orchestrator_for_organization" {
    provider = google-beta
    policy_orchestrator_id = "tf-test-test-po-org%{random_suffix}"
    organization_id = "%{org_id}"
    
    state = "STOPPED"
    action = "DELETE"
    description = "Updated description"
    
    orchestrated_resource {
        id = "tf-test-test-orchestrated-resource-org%{random_suffix}"
        os_policy_assignment_v1_payload {
            os_policies {
                id = "tf-test-test-os-policy-org%{random_suffix}"
                mode = "VALIDATION"
                resource_groups {
                    resources {
                        id = "resource-tf"
                        file {
                            content = "file-content-tf"
                            path = "file-path-tf-1"
                            state = "PRESENT"
                        }
                    }
                }
            }
            instance_filter {
                inventories {
                    os_short_name = "windows-10"
                }
            }
            rollout {
                disruption_budget {
                    percent = 100
                }
                min_wait_duration = "60s"
            }
        }
    }
}
`, context)
}
