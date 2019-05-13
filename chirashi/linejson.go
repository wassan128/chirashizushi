package chirashi

import (
    "encoding/json"
    "fmt"
)

type Container struct {
    Type string                 `json:"type"`
    Contents []ContainerContent `json:"contents"`
}
type ContainerContent struct {
    Type string `json:"type"`
    Body Body   `json:"body"`
}
type Body struct {
    Type string             `json:"type"`
    Layout string           `json:"layout"`
    Contents []BodyContent  `json:"contents"`
}
type BodyContent struct {
    Type string `json:"type"`
    Text string `json:"text"`
}

func GenerateMessage(items []Item) string {
    var container Container
    container.Type = "carousel"

    var contents []ContainerContent
    for _, item := range items {
        var content ContainerContent

        content.Type = "bubble"
        content.Body = Body{
            Type: "box",
            Layout: "vertical",
            Contents: []BodyContent{
                BodyContent{
                    Type: "text",
                    Text: fmt.Sprintf("%s(￥%d)", item.Name, item.Price),
                },
            },
        }
        contents = append(contents, content)
    }
    container.Contents = contents

    jsonBytes, _ := json.Marshal(&container)
    return string(jsonBytes)
}

