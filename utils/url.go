package utils

import (
	"fmt"
	"net/url"
	"strings"
)

func IsInternalLink(link string) bool {
	return strings.HasPrefix(link, "/")
}

func ExtractURL(rawURL string) []string {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil
	}
	path := parsedURL.Path
	path = strings.TrimPrefix(path, "/")
	return strings.Split(path, "/")
}
