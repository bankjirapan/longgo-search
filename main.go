package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	utils "longgo-search.com/utils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	url := "http://localhost:3000/docs3"

	var htmlContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.OuterHTML(`html`, &htmlContent),
	)
	if err != nil {
		log.Fatal(err)
	}
	_, links := utils.ParseHTML(strings.NewReader(htmlContent))

	// for _, content := range contents {
	// 	fmt.Println(content)
	// }
	// for text, href := range links {
	// 	fmt.Printf("%s: %s\n", text, href)
	// }

	// go to all links for scrap
	for _, href := range links {
		// fmt.Println(links)
		// var htmlContent string

		if utils.IsInternalLink(href) {
			// fmt.Println("Internal link is: " + href)
			url = "http://localhost:3000" + href
			fmt.Println(url)

			err := chromedp.Run(ctx,
				chromedp.Navigate(url),
				chromedp.WaitVisible(`body`, chromedp.ByQuery),
				chromedp.OuterHTML(`html`, &htmlContent),
			)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			contents, _ := utils.ParseHTML(strings.NewReader(htmlContent))
			fmt.Println(contents)

		} else {
			// /
		}
		fmt.Println()
	}

}
