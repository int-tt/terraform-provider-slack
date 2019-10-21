package slack

import (
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	slackapi "github.com/nlopes/slack"
)

//Provider is the root of terraform provider plugin
func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"SLACK_TOKEN",
				}, nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"slack_channel":        resourceChannel(),
			"slack_channel_invite": resourceChannelInvite(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"slack_user":            dataSourceUser(),
			"slack_user_with_email": dataSourceUserWithEmail(),
		},
		ConfigureFunc: providerConfigure,
		MetaReset:     nil,
	}

	return provider
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
