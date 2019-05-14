package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/line/line-bot-sdk-go/linebot"
    "github.com/wassan128/chirashizushi/chirashi"
)

func getMessage(shopId string) *linebot.CarouselContainer {
    shop := chirashi.Open(shopId)
    items := shop.GetTokubaiInfo()
    return chirashi.GenerateMessage(shop, items)
}

func main() {
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
                    container := getMessage(message.Text)
                    if _, err = bot.ReplyMessage(
                        event.ReplyToken,
                        linebot.NewFlexMessage("Chirashi submitted.", container),
                    ).Do(); err != nil {
                        log.Print(err)
                    }
                }
            }
        }
    })

    port := os.Getenv("PORT")
    if err := http.ListenAndServe(":" + port, nil); err != nil {
        log.Fatal(err)
    }
}
