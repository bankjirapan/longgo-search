package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	utils "longgo-search.com/utils"
)

func Crawler(urlStr string) {
	url, errParse := url.Parse(urlStr)
	if errParse != nil {
		log.Fatal(errParse)
	}
	baseURL := url.String()
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 5*time.Minute)
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
		var subURLContent string
		if utils.IsInternalLink(href) {

			url.Path = ""
			url.RawQuery = ""
			url.Fragment = ""

			subURL := url.String() + href
			fmt.Println("Crawling: ", subURL)
			err := chromedp.Run(ctx,
				chromedp.Navigate(subURL),
				chromedp.WaitVisible(`body`, chromedp.ByQuery),
				chromedp.Sleep(2*time.Second),
				chromedp.OuterHTML(`html`, &subURLContent),
			)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			contents, _ := utils.ParseHTML(strings.NewReader(subURLContent), "h1, h2, h3, h4, h5, h6, p")
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
	fmt.Println("Crawling done!")
}
