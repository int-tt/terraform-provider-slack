package slack

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	slack "github.com/nlopes/slack"
)

func resourceChannelInvite() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"channel_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_join": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
		Create:   resourceChannelInviteCreate,
		Read:     resourceChannelInviteRead,
		Update:   resourceChannelInviteUpdate,
		Delete:   resourceChannelInviteDelete,
		Importer: &schema.ResourceImporter{State: resourceChannelInviteImport},
	}
}

func resourceChannelInviteCreate(d *schema.ResourceData, meta interface{}) error {
	channelID, userID := getUserAndChannelID(d)
	_, err := meta.(*slack.Client).InviteUserToChannel(channelID, userID)
	if err != nil {
		return fmt.Errorf("faild to invite user to channel:%s", err.Error())
	}
	return resourceChannelInviteRead(d, meta)
}

func resourceChannelInviteRead(d *schema.ResourceData, meta interface{}) error {
	channelID, userID := getUserAndChannelID(d)
	channel, err := meta.(*slack.Client).GetChannelInfo(channelID)
	if err != nil {
		return fmt.Errorf("failed to read channel:%s", err.Error())
	}
	isJoin := "false"
	for _, member := range channel.Members {
		if member == userID {
			isJoin = "true"
			break
		}
	}

	d.SetId(isJoin)
	return nil
}

func resourceChannelInviteUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceChannelInviteRead(d, meta)
}

func resourceChannelInviteDelete(d *schema.ResourceData, meta interface{}) error {
	channelID, userID := getUserAndChannelID(d)
	err := meta.(*slack.Client).KickUserFromChannel(channelID, userID)
	if err != nil {
		return fmt.Errorf("failed to kick user to channel:%s", err.Error())
	}
	return resourceChannelInviteRead(d, meta)
}

func resourceChannelInviteImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := resourceChannelInviteRead(d, meta); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func getUserAndChannelID(d *schema.ResourceData) (string, string) {
	return d.Get("channel_id").(string), d.Get("user_id").(string)
}
