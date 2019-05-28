package chirashi

import (
    "log"
    "strconv"
    "strings"

    "github.com/PuerkitoBio/goquery"
    "github.com/wassan128/chirashizushi/util"
)

type Shop struct {
    Id string
    Name string
    chirashi *goquery.Document
}
type Item struct {
    ImageUrl string
    Name string
    Description string
    Price int
    Label string
}
var shopId string

func Open(id string) Shop {
    doc, err := goquery.NewDocument("https://tokubai.co.jp/" + id)
    if err != nil {
        log.Fatal(err)
    }

    name := util.Strip(doc.Find("a.shop_name").Text())

    var shop Shop
    shop.chirashi = doc
    shop.Id = id
    shop.Name = name

    return shop
}

func (shop Shop) GetTokubaiInfo() []Item {
    items := shop.chirashi.Find("a[id*=featured_product]")

    var tokubaiItems []Item
    items.Each(func(_ int, item *goquery.Selection) {
        var tokubaiItem Item
        // Thumb
        img := item.Find(".image")
        src, _ := img.Attr("data-src")
        if len(src) == 0 {
            src = "https://raw.githubusercontent.com/wassan128/chirashizushi/master/noimage.jpg"
        }
        tokubaiItem.ImageUrl = src

        // Label
        label := util.Strip(item.Find(".label_class").Text())
        if len(label) == 0 {
            comment := util.Strip(item.Find(".comment_wrapper").Text())
            if len(comment) == 0 {
                label = name
            } else {
                label = comment
            }
        }
        tokubaiItem.Label = label

        // Name
        name := util.Strip(item.Find(".name").Text())
        if len(name) == 0 {
            name = "(商品名未登録)"
        }
        tokubaiItem.Name = name

        // Description
        desc := util.Strip(item.Find(".price_unit_and_production_area").Text())
        if len(desc) == 0 {
            desc = name
        }
        tokubaiItem.Description = desc

        // Price
        price, err := strconv.Atoi(strings.Replace(item.Find(".number").Text(), ",", "", -1))
        tokubaiItem.Price = price

        if err == nil {
            tokubaiItems = append(tokubaiItems, tokubaiItem)
        }
    })

    return tokubaiItems
}

