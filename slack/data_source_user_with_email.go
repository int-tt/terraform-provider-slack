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
	userEmail := d.Get("email").(string)
	user, err := meta.(*slackapi.Client).GetUserByEmail(userEmail)
	if err != nil {
		return fmt.Errorf("faild to get user(%s): %s", userEmail, err.Error())
	}
	if err = setUserInfo(d, user); err != nil {
		return fmt.Errorf("faild to set user info:%s", err.Error())
	}
	return nil
}
