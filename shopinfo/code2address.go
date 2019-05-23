package shopinfo

import (
    "log"

    "github.com/PuerkitoBio/goquery"
)

func Code2Address(zipCode string) string {
    doc, err := goquery.NewDocument("https://api.nipponsoft.co.jp/zipcode/?search=" + zipCode)
    if err != nil {
        log.Fatal(err)
    }

    return doc.Find("#resultBox > div > table.list > tbody > tr:nth-child(2) > td:nth-child(5) > p.kanji").Text()
}

