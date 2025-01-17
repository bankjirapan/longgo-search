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
	"longgo-search.com/maid"
	utils "longgo-search.com/utils"
)

func Crawler(urlStr string) {
	maidService := &maid.MaidService{}
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
	var totalLinks int
	var jsonData []byte
	totalLinks = len(links)
	fmt.Println("Total links found: ", totalLinks)
	for _, href := range links {
		var subURLContent string
		if utils.IsInternalLink(href) {
			url.Path = ""
			url.RawQuery = ""
			url.Fragment = ""

			subURL := url.String() + href
			fmt.Println("Crawling: ", subURL)
			totalLinks--
			fmt.Println("Remaining links: ", totalLinks)
			err := chromedp.Run(ctx,
				chromedp.Navigate(subURL),
				chromedp.WaitVisible(`body`, chromedp.ByQuery),
				chromedp.Sleep(1*time.Second),
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
			fmt.Println("External link: ", href)
			// /
		}
	}
	if len(jsonData) > 0 {
		jsonData = jsonData[:len(jsonData)-1]
	}
	jsonData = append(jsonData, ']')
	jsonData = append([]byte{'['}, jsonData...)

	var jsonForMaid []maid.Content
	errCleaning := json.Unmarshal([]byte(jsonData), &jsonForMaid)
	if errCleaning != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	maidService.RemoveAllDuplicates(jsonForMaid)

	ioutil.WriteFile("data/web-data.json", jsonData, os.ModePerm)
	fmt.Println("Crawling done!")
}
