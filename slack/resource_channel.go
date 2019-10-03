package slack

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	slackapi "github.com/nlopes/slack"
)

func resourceChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceChannelCreate,
		Read:   resourceChannleRead,
		Update: resourceChannelUpdate,
		Delete: resourceChannelDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceChannelCreate(d *schema.ResourceData, meta interface{}) error {
	channel, err := meta.(*slackapi.Client).CreateChannel(d.Get("name").(string))
	if err != nil {
		return fmt.Errorf("failed to create channel: %s", err.Error())
	}
	d.SetId(channel.ID)

	return nil
}

func resourceChannleRead(d *schema.ResourceData, meta interface{}) error {
	if _, err := meta.(*slackapi.Client).GetChannelInfo(d.Id()); err != nil {
		return fmt.Errorf("failed to read channel: %s", err.Error())
	}
	return nil
}

func resourceChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	if _, err := meta.(*slackapi.Client).RenameChannel(d.Id(), d.Get("name").(string)); err != nil {
		return fmt.Errorf("faild to update channel: %s", err.Error())
	}
	return nil
}

func resourceChannelDelete(d *schema.ResourceData, meta interface{}) error {
	if err := meta.(*slackapi.Client).ArchiveChannel(d.Id()); err != nil {
		return fmt.Errorf("failed to archive channel: %s", err.Error())
	}
	return nil
}
