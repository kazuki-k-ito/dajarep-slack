# dajarep-slack
Slack bot in Go for detecting Japanese "DAJARE"

![image](https://github.com/allezsans/dajarep-slack/blob/image/image01.png)

## Installing
Just go get as below:

`go get github.com/allezsans/dajarep-slack`

### Description
- sentence.go
  - Return your "DAJARE" with sentence like the above image
- reaction.go
  - Add slack reaction your slack remark

### Dependencies
You have to also `go get` dependencies as below:

`go get github.com/nlopes/slack`

`go get github.com/kurehajime/dajarep`

## Usage
Set your Slack API token.

`export SLACK_BOT_TOKEN=YOUR_SLACK_TOKEN`

Build and execute binary as daemon

`go build sentence.go` or `go build reaction.go`

`./sentence` or `./reaction -r "your reaction name(default +1)"`
