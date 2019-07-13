package util

import (
    "encoding/json"
    "log"
    "net/http"
    "os"

    "golang.org/x/net/context"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/sheets/v4"
)

type Sheet struct {
    Service *sheets.Service
    Id string
}

func LoadSheets() *Sheet {
    var sheet Sheet

    config, err := google.ConfigFromJSON([]byte(
        os.Getenv("CREDENTIALS")),
        "https://www.googleapis.com/auth/spreadsheets.readonly",
    )
    if err != nil {
        log.Fatal(err)
    }

    client := getClient(config)
    sheet.Service, err = sheets.New(client)
    if err != nil {
        log.Fatal(err)
    }

    sheet.Id = os.Getenv("SHEET_ID_MASTER")

    return &sheet
}

func getClient(config *oauth2.Config) *http.Client {
    tok, _ := tokenFromEnv()
    return config.Client(context.Background(), tok)
}

func tokenFromEnv() (*oauth2.Token, error) {
    tok := &oauth2.Token{}
    err := json.Unmarshal([]byte(os.Getenv("TOKEN")), tok)
    return tok, err
}

func (sheet Sheet) ReadShopIds() map[string]string {
    readRange := "A2:B"
    res, err := sheet.Service.Spreadsheets.Values.Get(
        sheet.Id,
        readRange,
    ).Do()

    if err != nil {
        log.Fatal(err)
    }

    ret := map[string]string{}
    for _, row := range res.Values {
        shopId, _ := row[0].(string)
        shopName, _ := row[1].(string)
        ret[shopId] = shopName
    }

    return ret
}
