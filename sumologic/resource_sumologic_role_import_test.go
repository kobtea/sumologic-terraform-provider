// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Sumo Logic and manual
//     changes will be clobbered when the file is regenerated. Do not submit
//     changes to this file.
//
// ----------------------------------------------------------------------------\
package sumologic

import (
	"fmt"
	//"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestSumologicRole_import(t *testing.T) {
	var role Role
	testName := FieldsMap["Role"]["name"]
	testDescription := FieldsMap["Role"]["description"]
	testFilterPredicate := FieldsMap["Role"]["filterPredicate"]
	testCapabilities := []string{"\"" + FieldsMap["Role"]["capabilities"] + "\""}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRoleDestroy(role),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSumologicRoleConfigImported(testName, testDescription, testFilterPredicate, testCapabilities),
			},
			{
				ResourceName:      "sumologic_role.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSumologicRoleConfigImported(name string, description string, filterPredicate string, capabilities []string) string {
	return fmt.Sprintf(`
resource "sumologic_role" "foo" {
      name = "%s"
      description = "%s"
      filter_predicate = "%s"
      capabilities = %v
}
`, name, description, filterPredicate, capabilities)
}
