// maid cleaning your data and making it ready for search

package maid

type MaidInterface interface {
	RemoveAllDuplicates(jsonData []Content) []Content
}

type Content struct {
	Contents []string `json:"contents"`
	URL      string   `json:"url"`
}

type MaidService struct{}

func (m *MaidService) RemoveAllDuplicates(jsonData []Content) []Content {
	seen := make(map[string]bool)

	for i, item := range jsonData {
		uniqueContents := []string{}
		for _, content := range item.Contents {
			if !seen[content] {
				uniqueContents = append(uniqueContents, content)
				seen[content] = true
			}
		}
		jsonData[i].Contents = uniqueContents
	}

	return jsonData
}
