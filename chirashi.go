package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "strconv"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

type ShopInfo struct {
    Id string
    Name string
    Chirashi *goquery.Document
}
type Item struct {
    ImageUrl string
    Name string
    Price int
}

func openDoc() *goquery.Document {
    file, _ := ioutil.ReadFile("store2.html")
    reader := strings.NewReader(string(file))

    doc, err := goquery.NewDocumentFromReader(reader)
    if err != nil {
        log.Fatal(err)
    }

    return doc
}

func GetShopInfo(id string) ShopInfo {
    doc := openDoc()
    name := doc.Find(".shop_name").Text()

    var shop ShopInfo
    shop.Id = id
    shop.Name = name
    shop.Chirashi = doc

    return shop
}

func GetTokubaiInfo(shop ShopInfo) []Item {
    items := shop.Chirashi.Find("a[id^=office_featured_product]")

    var tokubaiItems []Item
    items.Each(func(_ int, item *goquery.Selection) {
		var tokubaiItem Item

        // Thumb
        img := item.Find(".image")
        src, _ := img.Attr("data-src")
        tokubaiItem.ImageUrl = src

        // Name
        name := strings.Trim(item.Find(".name").Text(), "\n")
        tokubaiItem.Name = name

        // Price
        price := item.Find(".number").Text()
        tokubaiItem.Price, _ = strconv.Atoi(strings.Replace(price, ",", "", -1))

        tokubaiItems = append(tokubaiItems, tokubaiItem)
    })

    return tokubaiItems
}

func main() {
    shop := GetShopInfo("")
    items := GetTokubaiInfo(shop)
    if len(items) == 0 {
        fmt.Printf("%sの特売情報はありません", shop.Name)
    }

    fmt.Printf("%sの特売情報です\n---\n", shop.Name)
    for _, item := range items {
        fmt.Printf("「%s」が%d円で売られています。", item.Name, item.Price)
        if len(item.ImageUrl) > 0 {
            fmt.Printf("(%s)", item.ImageUrl)
        }
        fmt.Println("\n---")
    }
}

