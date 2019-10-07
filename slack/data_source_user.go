package slack

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		SchemaVersion:      0,
		MigrateState:       nil,
		StateUpgraders:     nil,
		Create:             nil,
		Read:               nil,
		Update:             nil,
		Delete:             nil,
		Exists:             nil,
		CustomizeDiff:      nil,
		Importer:           nil,
		DeprecationMessage: "",
		Timeouts:           nil,
	}
}
