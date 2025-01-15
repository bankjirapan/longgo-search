package utils

import (
	"fmt"
	"io"
	"log"

	goquery "github.com/PuerkitoBio/goquery"
)

func ParseHTML(filePath io.Reader) ([]string, map[string]string) {
	doc, err := goquery.NewDocumentFromReader(filePath)
	if err != nil {
		fmt.Println("Error loading HTML file")
		log.Fatal(err)
	}

	var contents []string
	linkMap := make(map[string]string)
	doc.Find("title").Each(func(_ int, s *goquery.Selection) {
		title := s.Text()
		fmt.Println(title)
	})
	doc.Find("body").Each(func(_ int, s *goquery.Selection) {
		tagContents := make(map[string][]string)
		s.Find("p, h1, h2, h3, h4, h5, h6").Each(func(_ int, element *goquery.Selection) {
			tag := goquery.NodeName(element)
			text := element.Text()
			tagContents[tag] = append(tagContents[tag], text)
		})
		for tag, texts := range tagContents {
			contents = append(contents, fmt.Sprintf("%s: %v", tag, texts))
		}

		s.Find("a").Each(func(_ int, link *goquery.Selection) {
			href, exists := link.Attr("href")
			if exists {
				linkText := link.Text()
				linkMap[linkText] = href
			}
		})
	})

	return contents, linkMap
}
