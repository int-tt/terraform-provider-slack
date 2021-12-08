package slack

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	slackapi "github.com/slack-go/slack"
)

func dataSourceUserWithEmail() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"email": {
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
		Read: dataSourceUserWithEmailRead,
	}
}

func dataSourceUserWithEmailRead(d *schema.ResourceData, meta interface{}) error {
	user, err := meta.(*slackapi.Client).GetUserByEmail(d.Get("email").(string))
	if err != nil {
		return fmt.Errorf("faild to get user: %s", err.Error())
	}
	if err = setUserInfo(d, user); err != nil {
		return fmt.Errorf("faild to set user info:%s", err.Error())
	}
	return nil
}
