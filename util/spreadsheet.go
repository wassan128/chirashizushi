package util

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"

    "golang.org/x/net/context"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/sheets/v4"
)

func getClient(config *oauth2.Config) *http.Client {
    tok, _ := tokenFromEnv()
    return config.Client(context.Background(), tok)
}

func tokenFromEnv() (*oauth2.Token, error) {
    tok := &oauth2.Token{}
    err := json.Unmarshal([]byte(os.Getenv("TOKEN")), tok)
    return tok, err
}

func ReadShopIds() []string {
    config, err := google.ConfigFromJSON([]byte(
        os.Getenv("CREDENTIALS")),
        "https://www.googleapis.com/auth/spreadsheets.readonly",
    )
    if err != nil {
        log.Fatal(err)
    }

    client := getClient(config)

    srv, err := sheets.New(client)
    if err != nil {
        log.Fatal(err)
    }

    // TODO
    spreadsheetId := os.Getenv("SHEET_ID_MASTER")
    readRange := "A:A"
    res, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
    if err != nil {
        log.Fatal(err)
    }

    ret := []string{}
    for _, row := range res.Values {
        ret = append(ret, fmt.Sprintf("%s", row[0]))
    }

    return ret
}
