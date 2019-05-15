package shopinfo

import (
    "log"

    "github.com/PuerkitoBio/goquery"
    "github.com/wassan128/chirashizushi/util"
)

type ShopInfo struct {
    Name string
    Id string
}

func Search(zipCode string) []ShopInfo {
    doc, err := goquery.NewDocument("https://tokubai.co.jp/recommend?zip_code=" + zipCode)
    if err != nil {
        log.Fatal(err)
    }

    shops := doc.Find("label.shop")

    var shopInfos []ShopInfo
    shops.Each(func(_ int, shop *goquery.Selection) {
        var shopInfo ShopInfo
        shopInfo.Name = util.Strip(shop.Find(".shop_name").Text())
        extracted, _ := shop.Attr("id")
        shopInfo.Id = util.Strip(extracted[5:])
        shopInfos = append(shopInfos, shopInfo)
    })
    return shopInfos
}

