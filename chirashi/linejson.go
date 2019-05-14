package chirashi

import (
    "fmt"
    "sort"

    "github.com/line/line-bot-sdk-go/linebot"
)

func GenerateMessage(items []Item) *linebot.CarouselContainer {
    sort.Slice(items, func(i, j int) bool {
        return items[i].Price < items[j].Price
    })
    if len(items) > 10 {
        items = items[:10]
    }

    bubbles := []*linebot.BubbleContainer{}
    for _, item := range items {
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
                Contents: []linebot.FlexComponent{
                    &linebot.TextComponent{
                        Type: linebot.FlexComponentTypeText,
                        Size: linebot.FlexTextSizeTypeXl,
                        Weight: "bold",
                        Text: item.Name,
                    },
                    &linebot.TextComponent{
                        Type: linebot.FlexComponentTypeText,
                        Color: "#888888",
                        Text: item.Description,
                    },
                    &linebot.TextComponent{
                        Type: linebot.FlexComponentTypeText,
                        Size: linebot.FlexTextSizeType3xl,
                        Align: linebot.FlexComponentAlignTypeEnd,
                        Weight: "bold",
                        Color: "#ff0000",
                        Text: fmt.Sprintf("￥%d", item.Price),
                    },
                },
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

