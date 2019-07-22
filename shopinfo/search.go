package shopinfo

import (
    "log"

    "github.com/PuerkitoBio/goquery"
    "github.com/wassan128/chirashizushi/util"
)

type ShopAttr struct {
    Name string
    Id string
}
type ShopInfo struct {
    Category string
    Shops []ShopAttr
}

func Search(zipCode string) (string, []ShopInfo) {
    doc, err := goquery.NewDocument("https://tokubai.co.jp/recommend?zip_code=" + zipCode)
    if err != nil {
        log.Fatal(err)
    }

    var areaName string

    var shopInfos []ShopInfo
    categories := doc.Find(".nearest_shops_wrapper > div").
                        Not(".subscribe_recommended_shops, .title, .change_zip_code")
    categories.Each(func(_ int, category *goquery.Selection) {
        var shopInfo ShopInfo

        categoryName := category.Find("h2.business_category_name").Text()
        shopInfo.Category = categoryName

        var shopAttrs []ShopAttr
        shops := doc.Find("label.shop")
        areaName = Code2Address(zipCode)
        shops.Each(func(_ int, shop *goquery.Selection) {
            var shopAttr ShopAttr
            shopAttr.Name = util.Strip(shop.Find(".shop_name").Text())
            extracted, _ := shop.Attr("id")
            shopAttr.Id = util.Strip(extracted[5:])
            shopAttrs = append(shopAttrs, shopAttr)
        })
        shopInfo.Shops = shopAttrs

        shopInfos = append(shopInfos, shopInfo)
    })

    return areaName, shopInfos
}

