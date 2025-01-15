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
	Tags    []string `json:"tags"`
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
			var heading string
			plainContent := strings.ToLower(utils.RemoveHTMLTags(content))
			searchPhrase := strings.ToLower(phrase)
			headingH1 := utils.FindH1Index(doc.Contents)

			if headingH1 != -1 {
				heading = doc.Contents[headingH1]

			} else {
				heading = "No heading found"
			}

			if strings.Contains(plainContent, searchPhrase) {
				extractURL := utils.ExtractURL(doc.URL)
				matchingDocuments = append(matchingDocuments, Match{
					URL:     doc.URL,
					Heading: utils.RemovePrefix(heading),
					Text:    utils.RemovePrefixArray([]string{content}),
					Tags:    extractURL,
				})
			}
		}
	}
	return matchingDocuments

}
