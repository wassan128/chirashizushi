package main

import (
    "log"
    "net/http"
    "os"
    "strconv"

    "github.com/line/line-bot-sdk-go/linebot"
    "github.com/wassan128/chirashizushi/chirashi"
    "github.com/wassan128/chirashizushi/mybot"
)

func chirashiHandler(shopId string, replyToken string, bot *linebot.Client) {
    if _, err := strconv.Atoi(shopId); err != nil {
        bot.ReplyMessage(
            replyToken,
            linebot.NewTextMessage("お店IDは数字で送信してください"),
        ).Do()
        return
    }

    shop := chirashi.Open(shopId)
    if len(shop.Name) == 0 {
        bot.ReplyMessage(
            replyToken,
            linebot.NewTextMessage("そのようなIDの店舗は見つかりませんでした"),
        ).Do()
        return
    }

    items := shop.GetTokubaiInfo()
    if len(items) == 0 {
        bot.ReplyMessage(
            replyToken,
            linebot.NewTextMessage(shop.Name + "のチラシ情報です https://tokubai.co.jp/" + shop.Id),
        ).Do()
        return
    }

    container := mybot.GenerateChirashiMessage(items)
    bot.ReplyMessage(
        replyToken,
        linebot.NewTextMessage(shop.Name + "のチラシ情報です https://tokubai.co.jp/" + shop.Id),
        linebot.NewFlexMessage("チラシ情報です", container),
    ).Do()
}

func main() {
    bot, err := linebot.New(
        os.Getenv("CHANNEL_SECRET"),
        os.Getenv("CHANNEL_ACCESS_TOKEN"),
    )

    if err != nil {
        log.Fatal(err)
    }

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
                    chirashiHandler(message.Text, event.ReplyToken, bot)
                }
            }
        }
    })

    port := os.Getenv("PORT")
    if err := http.ListenAndServe(":" + port, nil); err != nil {
        log.Fatal(err)
    }
}
