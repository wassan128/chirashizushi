package main

import (
    "fmt"
    "encoding/json"

    "github.com/wassan128/chirashizushi/chirashi"
    "github.com/wassan128/chirashizushi/mybot"
)

func main() {
    shop := chirashi.Open("7192")
    items := shop.GetTokubaiInfo()
    container := mybot.GenerateMessage(items)

    jsonByte, _ := json.Marshal(container)
    fmt.Printf("%v\n", string(jsonByte))
}
