package slack

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	slackapi "github.com/nlopes/slack"
)

func setUserInfo(d *schema.ResourceData, user *slackapi.User) error {
	d.SetId(user.ID)
	_ = d.Set("team_id", user.TeamID)
	_ = d.Set("name", user.Name)
	_ = d.Set("real_name", user.RealName)
	return nil
}
