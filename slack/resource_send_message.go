package slack

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	slackapi "github.com/slack-go/slack"
)

func resourceSendMessage() *schema.Resource {
	return &schema.Resource{
		Create: resourceSendMessageCreate,
		Read:   resourceChannelRead,
		Update: resourceChannelUpdate,
		Delete: resourceChannelDelete,
		Schema: map[string]*schema.Schema{
			"channel_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"text": {
				Type:     schema.TypeString,
				Required: true,
			},
			"timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSendMessageCreate(d *schema.ResourceData, meta interface{}) error {
	// https://github.com/nlopes/slack/blob/v0.6.0/chat.go#L283
	channel, timestamp, _, err := meta.(*slackapi.Client).SendMessage(
		d.Get("channel_id").(string), slackapi.MsgOptionText(d.Get("text").(string), false),
	)
	if err != nil {
		return fmt.Errorf("failed to send message: %s", err.Error())
	}
	d.SetId(fmt.Sprintf("%s:%s", channel, timestamp))
	d.Set("timestamp", timestamp)

	return resourceSendMessageRead(d, meta)
}

func resourceSendMessageRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceSendMessageUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceSendMessageDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
