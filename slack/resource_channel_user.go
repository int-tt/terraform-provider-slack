package slack

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceChannelUser() *schema.Resource {
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
		},
		SchemaVersion:      0,
		MigrateState:       nil,
		StateUpgraders:     nil,
		Create:             resourceChannelUserCreate,
		Read:               resourceChannelUserRead,
		Update:             resourceChannelUserUpdate,
		Delete:             resourceChannelUserDelete,
		Exists:             nil,
		CustomizeDiff:      nil,
		Importer:           &schema.ResourceImporter{State: resourceChannelUserImport},
		DeprecationMessage: "",
		Timeouts:           nil,
	}
}

func resourceChannelUserCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceChannelUserRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceChannelUserUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceChannelUserDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceChannelUserImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{d}, nil
}
