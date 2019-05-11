package main

import (
    "log"
    "strings"
    "time"

    "github.com/sclevine/agouti"
)

func main() {
    // Initialize headless chrome driver
    /*
    driver := agouti.ChromeDriver(
        agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu"}))
    */
    driver := agouti.ChromeDriver()

    if err := driver.Start(); err != nil {
        log.Fatalf("driver start: %v", err)
    }
    defer driver.Stop()

    page, err := driver.NewPage(agouti.Browser("chrome"))
    if err != nil {
        log.Fatalf("new page: %v", err)
    }

    if page.Navigate("http://localhost:8000"); err != nil {
        log.Fatalf("navigte page: %v", err)
    }

    body := page.Find("body")
    body.SendKeys(strings.Repeat("\uE015", 5)) // keydown

    page.Screenshot("screenshot.png")
}
