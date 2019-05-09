package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

func main() {
    file, _ := ioutil.ReadFile("store2.html")
    reader := strings.NewReader(string(file))

    doc, err := goquery.NewDocumentFromReader(reader)
    if err != nil {
        panic(err)
    }

    items := doc.Find("a[id^=office_featured_product]")
    if items.Size() <= 0 {
        fmt.Printf("no chirashi information")
        os.Exit(-1)
    }

    items.Each(func(_ int, item *goquery.Selection) {
        // Thumb
        img := item.Find(".image")
        src, _ := img.Attr("data-src")
        fmt.Println(src)

        // Name
        name := strings.Trim(item.Find(".name").Text(), "\n")
        fmt.Println(name)

        // Price
        price := item.Find(".number").Text()
        fmt.Println(price)

        fmt.Println(strings.Repeat("-", 20))
    })
}
