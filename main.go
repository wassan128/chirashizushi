package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/line/line-bot-sdk-go/linebot"
)

func main() {
    //shop := chirashi.Open("7192")
    //fmt.Printf("Shop: [%+v]\n", shop)

    //items := shop.GetTokubaiInfo()
    //fmt.Printf("Items: [%+v]\n", items)

    //json := chirashi.GenerateMessage(items)
    //fmt.Printf("%+v\n", json)

    bot, err := linebot.New(
        os.Getenv("CHANNEL_SECRET"),
        os.Getenv("CHANNEL_ACCESS_TOKEN"),
    )

    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "Welcome to ようこそ ChirashiZushi world")
    })
    http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
        events, err := bot.ParseRequest(req)
        if err != nil {
            if err == linebot.ErrInvalidSignature {
                w.WriteHeader(400)
            } else {
                w.WriteHeader(500)
            }
            return
        }
        for _, event := range events {
            if event.Type == linebot.EventTypeMessage {
                fmt.Printf("%+v\n", event.Message)
                switch message := event.Message.(type) {
                case *linebot.TextMessage:
                    if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
                        log.Print(err)
                    }
                }
            }
        }
    })

    if err := http.ListenAndServe(":5000", nil); err != nil {
        log.Fatal(err)
    }
}
