package utils

import "strings"

func IsInternalLink(link string) bool {
	return strings.HasPrefix(link, "/")
}
