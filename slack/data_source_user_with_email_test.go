package slack

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccDataSourceUserWithEmail(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUserWithEmailConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceUserWithEmailCheck("data.slack_user_with_email.foo"),
				),
			},
		},
	})
}

func testAccDataSourceUserWithEmailConfig() string {
	return fmt.Sprintf(`
data "slack_user_with_email" "foo" {
  email = "kattu0426@gmail.com"
}
`)
}

func testAccDataSourceUserWithEmailCheck(dataSourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[dataSourceName]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", dataSourceName)
		}
		dsAttr := ds.Primary.Attributes
		userTestSExcept := map[string]string{
			"id":        "UP0H1PAN8",
			"name":      "kattu0426",
			"real_name": "kattu0426",
		}
		keys := []string{
			"id",
			"name",
			"real_name",
		}
		for _, k := range keys {
			if userTestSExcept[k] != dsAttr[k] {
				return fmt.Errorf("%s is %s; want %s", k, dsAttr[k], userTestSExcept[k])
			}
		}
		return nil
	}
}
