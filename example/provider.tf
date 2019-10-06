provider "slack" {
  token = ""
}

resource "slack_channel" "test" {
  name = "test-channel"
}