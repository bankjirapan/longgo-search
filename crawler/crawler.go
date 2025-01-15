package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	utils "longgo-search.com/utils"
)

func Crawler() {
	var baseURL = "http://localhost:3000/docs3"
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	var htmlContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(baseURL),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.OuterHTML(`html`, &htmlContent),
	)
	if err != nil {
		log.Fatal(err)
	}
	_, links := utils.ParseHTML(strings.NewReader(htmlContent), "")
	var jsonData []byte
	for _, href := range links {
		if utils.IsInternalLink(href) {
			subURL := "http://localhost:3000" + href
			fmt.Println(subURL)
			err := chromedp.Run(ctx,
				chromedp.Navigate(subURL),
				chromedp.WaitVisible(`body`, chromedp.ByQuery),
				chromedp.OuterHTML(`html`, &htmlContent),
			)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			contents, _ := utils.ParseHTML(strings.NewReader(htmlContent), "p, h1")
			result := map[string]interface{}{
				"url":      subURL,
				"contents": contents,
			}
			jsonResult, err := json.Marshal(result)
			if err != nil {
				fmt.Println("Error marshalling to JSON: ", err)
			}
			jsonData = append(jsonData, jsonResult...)
			jsonData = append(jsonData, ',')

		} else {
			// /
		}
	}
	if len(jsonData) > 0 {
		jsonData = jsonData[:len(jsonData)-1]
	}
	jsonData = append(jsonData, ']')
	jsonData = append([]byte{'['}, jsonData...)
	ioutil.WriteFile("data/web-data.json", jsonData, os.ModePerm)
}
