# dajarep-slack
Slack bot in Go for detecting Japanese "DAJARE"

## Installing
Just go get as below:

`go get github.com/allezsans/dajarep-slack`

### Dependencies
You have to also `go get` dependencies as below:

`go get github.com/nlopes/slack`

`go get github.com/kurehajime/dajarep`

## Usage
Set your Slack API token.

`export SLACK_BOT_TOKEN=YOUR_SLACK_TOKEN`

Build and execute binary as daemon

`go build main.go`

`./main`
