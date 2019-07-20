package mybot

import (
    "fmt"
    "sort"

    "github.com/line/line-bot-sdk-go/linebot"
    "github.com/wassan128/chirashizushi/chirashi"
    "github.com/wassan128/chirashizushi/shopinfo"
)

func GenerateChirashiMessage(items []chirashi.Item) *linebot.CarouselContainer {
    sort.Slice(items, func(i, j int) bool {
        return items[i].Price < items[j].Price
    })
    if len(items) > 10 {
        items = items[:10]
    }

    bubbles := []*linebot.BubbleContainer{}
    for _, item := range items {
        var components []linebot.FlexComponent
        components = append(components,
            &linebot.TextComponent{
                Type: linebot.FlexComponentTypeText,
                Size: linebot.FlexTextSizeTypeLg,
                Weight: "bold",
                Text: item.Name,
            },
            &linebot.TextComponent{
                Type: linebot.FlexComponentTypeText,
                Size: linebot.FlexTextSizeTypeSm,
                Color: "#888888",
                Text: item.Description,
            },
            &linebot.TextComponent{
                Type: linebot.FlexComponentTypeText,
                Size: linebot.FlexTextSizeTypeXxl,
                Align: linebot.FlexComponentAlignTypeEnd,
                Weight: "bold",
                Color: "#ff3366",
                Text: fmt.Sprintf("￥%d", item.Price),
            },
        )
        if len(item.Label) > 0 {
           components = append(components,
               &linebot.SeparatorComponent{
                    Type: linebot.FlexComponentTypeSeparator,
                    Color: "#eeeeee",
                },
                &linebot.TextComponent{
                    Type: linebot.FlexComponentTypeText,
                    Size: linebot.FlexTextSizeTypeXs,
                    Margin: linebot.FlexComponentMarginTypeMd,
                    Align: linebot.FlexComponentAlignTypeEnd,
                    Color: "#666666",
                    Text: item.Label,
                },
            )
        }
        bubble := &linebot.BubbleContainer{
            Type: linebot.FlexContainerTypeBubble,
            Hero: &linebot.ImageComponent{
                Type: linebot.FlexComponentTypeImage,
                URL: item.ImageUrl,
                Size: linebot.FlexImageSizeTypeFull,
                AspectRatio: linebot.FlexImageAspectRatioType1to1,
                AspectMode: linebot.FlexImageAspectModeTypeCover,
            },
            Body: &linebot.BoxComponent{
                Type: linebot.FlexComponentTypeBox,
                Layout: linebot.FlexBoxLayoutTypeVertical,
                Contents: components,
            },
        }
        bubbles = append(bubbles, bubble)
    }
    contents := &linebot.CarouselContainer{
        Type: linebot.FlexContainerTypeCarousel,
        Contents: bubbles,
    }

    return contents
}

func GenerateShopInfoMessage(shopinfos []shopinfo.ShopInfo) string {
    shopinfoMsg := ""
    for _, shopinfo := range shopinfos {
        shopinfoMsg += fmt.Sprintf("【%s】%s\n", shopinfo.Id, shopinfo.Name)
    }
    return shopinfoMsg
}
