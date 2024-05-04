package slack

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	slack "github.com/slack-go/slack"
	slackapi "github.com/slack-go/slack"
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
	for {
		_, err := meta.(*slack.Client).InviteUsersToConversation(channelID, userID)
		if err != nil {
			if e, ok := err.(*slackapi.RateLimitedError); ok {
				time.Sleep(e.RetryAfter)
				continue
			}
			return fmt.Errorf("faild to invite user to channel:%s", err.Error())
		}
		break
	}
	return resourceChannelInviteRead(d, meta)
}

func resourceChannelInviteRead(d *schema.ResourceData, meta interface{}) error {
	channelID, userID := getUserAndChannelID(d)
	for {
		channel, err := meta.(*slack.Client).GetConversationInfo(channelID, false)
		if err != nil {
			if e, ok := err.(*slackapi.RateLimitedError); ok {
				time.Sleep(e.RetryAfter)
				continue
			}
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
		break
	}
	return nil
}

func resourceChannelInviteUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceChannelInviteRead(d, meta)
}

func resourceChannelInviteDelete(d *schema.ResourceData, meta interface{}) error {
	channelID, userID := getUserAndChannelID(d)
	for {
		err := meta.(*slack.Client).KickUserFromConversation(channelID, userID)
		if err != nil {
			if e, ok := err.(*slackapi.RateLimitedError); ok {
				time.Sleep(e.RetryAfter)
				continue
			}

			return fmt.Errorf("failed to kick user to channel:%s", err.Error())
		}
		break
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
