package slack

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	slackapi "github.com/nlopes/slack"
)
func resourceSendMessage() *schema.Resource {
	return &schema.Resource{
		Create: resourceSendMessageCreate,
		Read:   resourceChannelRead,
		Update: resourceChannelUpdate,
		Delete: resourceChannelDelete,
		Schema: map[string]*schema.Schema{
			"channel_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"text": {
				Type: schema.TypeString,
				Required: true,
			},
			"timestamp": {
				Type: schema.TypeString,
				Computed: true,
			}
		},
	}
}

func resourceSendMessageCreate(d *schema.ResouceData, meta interface{}) error{
	// https://github.com/nlopes/slack/blob/v0.6.0/chat.go#L283
 	channel, timestamp, text, err := meta(*slackapi.Client).SendMessage(
		d.Get("channel_id"), slackapi.MsgOptionText(d.Get("text"))
	)
	if err != nil{
		return fmt.Errorf("failed to send message: %s", err.Error())
	}
	d.SetID(fmt.Sprintf("%s:%s",channel, timestamp))
	d.Set("timestamp", timestamp)
	
	return resourceSendMessageRead(d, meta)
}
func resourceSendMessageRead(d *schema.ResouceData, meta interface{}) error {
}
func resourceSendMessageUpdate(d *schema.ResourceData, meta interface{}) error{
}
func resourceSendMessageDelete(d *schema.ResouceData, meta interface{}) error {
}