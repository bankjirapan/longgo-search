package utils

func StringToJson(str string) string {
	// check last string is comma or not
	if str[len(str)-1:] == "," {
		str = str[:len(str)-1]
	}
	return "{" + str + "}"
}
