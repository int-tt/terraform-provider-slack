package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/int-tt/terraform-provider-slack/slack"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: slack.Provider,
	})
}
