package chirashi

import (
    "log"
    "strconv"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

type Shop struct {
    Id string
    Name string
    Chirashi *goquery.Document
}
type Item struct {
    ImageUrl string
    Name string
    Price int
}
var shopId string

func Open(id string) Shop {
    doc, err := goquery.NewDocument("https://tokubai.co.jp/" + id)
    if err != nil {
        log.Fatal(err)
    }

    name := Strip(doc.Find("a.shop_name").Text())

    var shop Shop
    shop.Chirashi = doc
    shop.Id = id
    shop.Name = name

    return shop
}

func (shop Shop) GetTokubaiInfo() []Item {
    items := shop.Chirashi.Find("a[id^=office_featured_product]")

    var tokubaiItems []Item
    items.Each(func(_ int, item *goquery.Selection) {
        var tokubaiItem Item
        // Thumb
        img := item.Find(".image")
        src, _ := img.Attr("data-src")
        tokubaiItem.ImageUrl = src

        // Name
        name := Strip(item.Find(".name").Text())
        tokubaiItem.Name = name

        // Price
        price := item.Find(".number").Text()
        tokubaiItem.Price, _ = strconv.Atoi(strings.Replace(price, ",", "", -1))

        tokubaiItems = append(tokubaiItems, tokubaiItem)
    })

    return tokubaiItems
}
