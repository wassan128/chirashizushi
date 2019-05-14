package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"

    "github.com/line/line-bot-sdk-go/linebot"
    "github.com/wassan128/chirashizushi/chirashi"
)

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
                switch message := event.Message.(type) {
                case *linebot.TextMessage:
                    if _, err := strconv.Atoi(message.Text); err != nil {
                        bot.ReplyMessage(
                            event.ReplyToken,
                            linebot.NewTextMessage("お店IDは数字で送信してください"),
                        ).Do()
                    }

                    shop := chirashi.Open(message.Text)
                    fmt.Printf("%+v\n", shop)
                    items := shop.GetTokubaiInfo()
                    fmt.Printf("%+v\n", items)

                    if len(items) == 0 {
                        bot.ReplyMessage(
                            event.ReplyToken,
                            linebot.NewTextMessage("チラシが見つかりませんでした"),
                        ).Do()
                        fmt.Printf("no chirashi")
                    } else {
                        container := chirashi.GenerateMessage(shop, items)
                        bot.ReplyMessage(
                            event.ReplyToken,
                            linebot.NewFlexMessage("チラシ情報です", container),
                        ).Do()
                        fmt.Printf("%+v\n", container)
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
