package airbyte

//
//import (
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
//	"regexp"
//	"testing"
//)
//
//var resourceTests = []struct {
//	resource   string
//	config     string
//	attrChecks map[string]string
//}{
//	{
//		"grafana_data_source.testdata",
//		`
//resource "airbyte_source" "source_1" {
//  name = "new-workspace"
//  workspace_id = airbyte_workspace.default_workspace.id
//  source_definition_id = "6371b14b-bc68-4236-bfbd-468e8df8e968"
//
//  connection_configuration = {
//    "pokemon_name": "ditto"
//  }
//}
//`,
//		map[string]string{
//			"type":                "testdata",
//			"name":                "testdata",
//			"access_mode":         "direct",
//			"basic_auth_enabled":  "true",
//			"basic_auth_password": "ba_password",
//			"basic_auth_username": "ba_username",
//			"database_name":       "db_name",
//			"is_default":          "true",
//			"url":                 "http://acc-test.invalid/",
//			"username":            "user",
//			"password":            "pass",
//		},
//	},
//}
//
//func TestAccDataSource_basic(t *testing.T) {
//	//var dataSource gapi.DataSource
//
//	// Iterate over the provided configurations for datasources
//	for _, test := range resourceTests {
//		// Always check that the resource was created and that `id` is a number
//		checks := []resource.TestCheckFunc{
//			//testAccDataSourceCheckExists(test.resource, &dataSource),
//			resource.TestMatchResourceAttr(
//				test.resource,
//				"id",
//				regexp.MustCompile(`\d+`),
//			),
//		}
//
//		// Add custom checks for specified attribute values
//		for attr, value := range test.attrChecks {
//			checks = append(checks, resource.TestCheckResourceAttr(
//				test.resource,
//				attr,
//				value,
//			))
//		}
//
//		resource.Test(t, resource.TestCase{
//			//PreCheck:          func() { testAccPreCheck(t) },
//			//ProviderFactories: testAccProviderFactories,
//			//CheckDestroy:      testAccDataSourceCheckDestroy(&dataSource),
//			Steps: []resource.TestStep{
//				{
//					Config: test.config,
//					Check: resource.ComposeAggregateTestCheckFunc(
//						checks...,
//					),
//				},
//			},
//		})
//	}
//}

//
//func testAccDataSourceCheckExists(rn string, dataSource *gapi.DataSource) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		rs, ok := s.RootModule().Resources[rn]
//		if !ok {
//			return fmt.Errorf("resource not found: %s", rn)
//		}
//
//		if rs.Primary.ID == "" {
//			return fmt.Errorf("resource id not set")
//		}
//
//		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
//		if err != nil {
//			return fmt.Errorf("resource id is malformed")
//		}
//
//		client := testAccProvider.Meta().(*client)
//		gotDataSource, err := client.DataSource(id)
//		if err != nil {
//			return fmt.Errorf("error getting data source: %s", err)
//		}
//
//		*dataSource = *gotDataSource
//
//		return nil
//	}
//}
//
//func testAccDataSourceCheckDestroy(source airbyte_sdk.SourceRead) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		client := testAccProvider.Meta().(*client)
//		_, err := client.DataSource(dataSource.ID)
//		if err == nil {
//			return fmt.Errorf("data source still exists")
//		}
//		return nil
//	}
//}
