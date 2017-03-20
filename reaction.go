package main

import (
    "flag"
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

    reaction := flag.String("r", "+1", "set your slack reaction")
    flag.Parse()

    rtm := api.NewRTM()
    go rtm.ManageConnection()

Loop:
    for {
       select {
        case msg := <-rtm.IncomingEvents:
            switch ev := msg.Data.(type) {
            case *slack.MessageEvent:
                msgRef := slack.NewRefToMessage(ev.Channel, ev.Timestamp)
                if res := hasDajare(ev.Text); res {
                    if err := api.AddReaction(*reaction, msgRef); err != nil {
                        fmt.Printf("Error adding reaction: %s\n", err)
                        return
                    }
                }
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

func hasDajare(msg string) bool {
    d, _ := dajarep.Dajarep(msg)
    return len(d) > 0
}
