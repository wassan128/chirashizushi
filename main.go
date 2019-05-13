package main

import (
    "fmt"

    "./chirashi"
)

func main() {
    shop := chirashi.Open("7192")
    fmt.Printf("Shop: [%+v]\n", shop)

    items := shop.GetTokubaiInfo()
    fmt.Printf("Items: [%+v]\n", items)

    json := chirashi.GenerateMessage(items)
    fmt.Printf("%+v\n", json)
}
