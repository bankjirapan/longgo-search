package utils

import (
	"fmt"
	"io"
	"log"

	goquery "github.com/PuerkitoBio/goquery"
)

func ParseHTML(filePath io.Reader, htmlTag string) ([]string, map[string]string) {
	doc, err := goquery.NewDocumentFromReader(filePath)
	if err != nil {
		fmt.Println("Error loading HTML file")
		log.Fatal(err)
	}

	var contents []string
	linkMap := make(map[string]string)
	doc.Find("body").Each(func(_ int, s *goquery.Selection) {
		if htmlTag == "" {
			text := s.Text()
			if text != "" {
				contents = append(contents, text)
			}
		} else {
			if htmlTag == "h1" {
				s.Find("h1").Each(func(_ int, element *goquery.Selection) {
					text := element.Text()
					if text != "" {
						contents = append(contents, "Header: "+text)
					}
				})
			} else {
				s.Find(htmlTag).Each(func(_ int, element *goquery.Selection) {
					text := element.Text()
					tagName := element.Nodes[0].Data
					if text != "" {
						contents = append(contents, tagName+": "+text)
					}
				})
			}
		}

		s.Find("a").Each(func(_ int, link *goquery.Selection) {
			href, exists := link.Attr("href")
			if exists {
				linkText := link.Text()
				if linkText != "" {
					linkMap[linkText] = href
				}
			}
		})
	})
	contents = removeDuplicates(contents)

	return contents, linkMap
}

func removeDuplicates(slice []string) []string {
	seen := make(map[string]struct{})
	var result []string
	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}
