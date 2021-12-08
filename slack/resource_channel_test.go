package slack

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"

	slackapi "github.com/slack-go/slack"
)

func TestAccSlackChannel_Basic(t *testing.T) {

	channelName := fmt.Sprintf("tf_test_%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckChannelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccChannelResourceConfig(channelName),
			},
			{
				ResourceName:      "slack_channel.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Update test
				Config: testAccChannelResourceConfig(fmt.Sprintf("%s_updated", channelName)),
			},
			{
				ResourceName:      "slack_channel.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccChannelResourceConfig(channelName string) string {
	return fmt.Sprintf(`
resource "slack_channel" "test" {
	name = "%s"
}
`, channelName)
}

func testAccCheckChannelDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "slack_channel" {
			continue
		}
		config := testAccProvider.Meta().(*slackapi.Client)
		channel, err := config.GetConversationInfo(rs.Primary.ID, false)
		if err != nil {
			return err
		}
		if !channel.IsArchived {
			return fmt.Errorf("channel still present")
		}

	}

	return nil
}
