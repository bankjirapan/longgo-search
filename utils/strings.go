package utils

func MergeStringSlices(slices ...[]string) []string {
	var mergedSlice []string
	for _, slice := range slices {
		mergedSlice = append(mergedSlice, slice...)
	}
	return mergedSlice
}

func RemoveHTMLTags(content string) string {
	return content
}
