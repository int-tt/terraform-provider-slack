provider "slack" {
  token = ""
}

data "slack_user" "user" {
  id = "UPFA0RTA4"
}

resource "slack_channel" "test" {
  name = "test-channel-invite-test-06"
}

resource "slack_channel_invite" "invite" {
  channel_id = slack_channel.test.id
  user_id = data.slack_user.user.id
}