provider "slack" {
  token = "SLACK_TOKEN"
}

resource "slack_channel" "test" {
  name = "test-channel"
}