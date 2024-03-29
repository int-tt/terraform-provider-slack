package slack

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	slackapi "github.com/slack-go/slack"
)

func resourceChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceChannelCreate,
		Read:   resourceChannelRead,
		Update: resourceChannelUpdate,
		Delete: resourceChannelDelete,
		Importer: &schema.ResourceImporter{
			State: resourceChannelImport,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
		},
	}
}

func resourceChannelCreate(d *schema.ResourceData, meta interface{}) error {
	isPrivate := d.Get("private").(bool)
	channelName := d.Get("name").(string)
	channel, err := meta.(*slackapi.Client).CreateConversation(channelName, isPrivate)
	if err != nil {
		return fmt.Errorf("failed to create channel(%s): %s", channelName, err.Error())
	}
	d.SetId(channel.ID)

	return resourceChannelRead(d, meta)
}

func resourceChannelRead(d *schema.ResourceData, meta interface{}) error {
	channel, err := meta.(*slackapi.Client).GetConversationInfo(d.Id(), false)
	if err != nil {
		return fmt.Errorf("failed to read channel: %s", err.Error())
	}
	if channel.IsArchived {
		return fmt.Errorf("failed to channel for archived")
	}
	_ = d.Set("name", channel.Name)

	return nil
}

func resourceChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	if _, err := meta.(*slackapi.Client).RenameConversation(d.Id(), d.Get("name").(string)); err != nil {
		return fmt.Errorf("failed to update channel: %s", err.Error())
	}

	return resourceChannelRead(d, meta)
}

func resourceChannelDelete(d *schema.ResourceData, meta interface{}) error {
	if err := meta.(*slackapi.Client).ArchiveConversation(d.Id()); err != nil {
		return fmt.Errorf("failed to archive channel: %s", err.Error())
	}
	return nil
}

func resourceChannelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := resourceChannelRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
