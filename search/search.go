package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"longgo-search.com/utils"
)

type Document struct {
	Contents []string `json:"contents"`
	URL      string   `json:"url"`
}

type Searcher interface {
	Search(phrase string) string
}

type Match struct {
	URL     string   `json:"url"`
	Heading string   `json:"heading"`
	Text    []string `json:"text"`
}

func Search(phrase string) []Match {

	data, err := ioutil.ReadFile("data/web-data.json")

	var documents []Document
	var matchingDocuments []Match
	err = json.Unmarshal([]byte(data), &documents)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}
	for _, doc := range documents {
		for _, content := range doc.Contents {
			plainContent := strings.ToLower(utils.RemoveHTMLTags(content))
			searchPhrase := strings.ToLower(phrase)
			headingH1 := utils.FindH1Index(doc.Contents)

			if strings.Contains(plainContent, searchPhrase) {
				matchingDocuments = append(matchingDocuments, Match{
					URL:     doc.URL,
					Heading: utils.RemovePrefix(doc.Contents[headingH1]),
					Text:    utils.RemovePrefixArray([]string{content}),
				})
			}
		}
	}
	return matchingDocuments

}
