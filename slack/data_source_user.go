package slack

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	slackapi "github.com/nlopes/slack"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"team_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"real_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Read: dataSourceUserRead,
	}
}

func dataSourceUserRead(d *schema.ResourceData, meta interface{}) error {
	user, err := meta.(*slackapi.Client).GetUserInfo(d.Get("id").(string))
	if err != nil {
		return fmt.Errorf("faild to get user: %s", err.Error())
	}
	d.SetId(user.ID)
	d.Set("team_id", user.TeamID)
	d.Set("name", user.Name)
	d.Set("real_name", user.RealName)
	return nil
}
