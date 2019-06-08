package main

import (
    "log"
    "net/http"
    "os"
    "strings"
    "strconv"

    "github.com/line/line-bot-sdk-go/linebot"
    "github.com/wassan128/chirashizushi/chirashi"
    "github.com/wassan128/chirashizushi/mybot"
    "github.com/wassan128/chirashizushi/shopinfo"
    "github.com/wassan128/chirashizushi/util"
)

func chirashiHandler(shopId, replyToken string, bot *linebot.Client) {
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

func menuHandler(text, replyToken string, bot *linebot.Client) {
    shopIds := util.ReadShopIds()

    if cmd := strings.Split(text, " "); len(cmd) == 1 {
        shopButtons := []*linebot.QuickReplyButton{}
        for shopId, shopName := range shopIds {
            shopButtons = append(shopButtons,
                linebot.NewQuickReplyButton("", linebot.NewMessageAction(shopName, shopId)))
        }
        shopButtons = append(shopButtons,
            linebot.NewQuickReplyButton("", linebot.NewLocationAction("現在地から探す")))

        bot.ReplyMessage(
            replyToken,
            linebot.NewTextMessage("アクションを選択してください").WithQuickReplies(
                &linebot.QuickReplyItems{
                    Items: shopButtons,
                },
            ),
        ).Do()
    } else {
        switch cmd[1] {
        case "セット":
            bot.ReplyMessage(
                replyToken,
                linebot.NewTextMessage("セットが指定されました"),
            ).Do()

        case "リセット":
            bot.ReplyMessage(
                replyToken,
                linebot.NewTextMessage("リセットが指定されました"),
            ).Do()

        case "テスト":

        default:
            bot.ReplyMessage(
                replyToken,
                linebot.NewTextMessage("不明なサブコマンドです"),
            ).Do()
        }
    }
}

func shopinfoHandler(zipCode, replyToken string, bot *linebot.Client) {
    if code := strings.Split(zipCode, "-"); len(code[0]) != 3 || len(code[1]) != 4  {
        bot.ReplyMessage(
            replyToken,
            linebot.NewTextMessage("郵便番号が不正です"),
        ).Do()
        return
    }

    areaName, shopinfos := shopinfo.Search(zipCode)
    if areaName == "404" || len(shopinfos) == 0 {
        bot.ReplyMessage(
            replyToken,
            linebot.NewTextMessage("ご指定の地域もしくは店舗が見つかりませんでした"),
        ).Do()
        return
    }

    container := mybot.GenerateShopInfoMessage(shopinfos)
    bot.ReplyMessage(
        replyToken,
        linebot.NewTextMessage(areaName + "(〒" + zipCode + ")の店舗リストです\n" + container),
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
                    text := message.Text
                    if strings.Contains(text, "チラシ") {
                        menuHandler(text, event.ReplyToken, bot)
                    } else if strings.Contains(text, "-") {
                        shopinfoHandler(text, event.ReplyToken, bot)
                    } else {
                        chirashiHandler(text, event.ReplyToken, bot)
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
