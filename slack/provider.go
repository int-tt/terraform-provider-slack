package slack

import (
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform/helper/logging"
	"github.com/hashicorp/terraform/helper/schema"
	slackapi "github.com/nlopes/slack"
)

//Provider is the root of terraform provider plugin
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"slack_channel": resourceChannel(),
		},
		DataSourcesMap: nil,
		ConfigureFunc:  providerConfigure,
		MetaReset:      nil,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	c := cleanhttp.DefaultClient()
	c.Transport = logging.NewTransport("Slack", c.Transport)
	client := slackapi.New(
		d.Get("token").(string),
		slackapi.OptionHTTPClient(c),
	)
	return client, nil
}
