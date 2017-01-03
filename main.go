package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/kurehajime/dajarep"
	"github.com/nlopes/slack"
)

func main() {
	token := os.Getenv("SLACK_BOT_TOKEN")
	api := slack.New(token)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				res := buildResponse(ev.Text)
				rtm.SendMessage(rtm.NewOutgoingMessage(res, ev.Channel))
			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop
			default:
				// Ignore other events..
				// fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}
}

func buildResponse(msg string) string {
	d, _ := dajarep.Dajarep(msg)
	var buffer bytes.Buffer
	for _, v := range d {
		buffer.WriteString("えっ?ダジャレっすか？？\n> ")
		buffer.WriteString(v)
		buffer.WriteString("\n")
	}
	return buffer.String()
}
