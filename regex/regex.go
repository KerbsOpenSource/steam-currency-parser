package regex

import (
	"regexp"
)

func OnlyInt(text string) string {
	reg := regexp.MustCompile("[^0-9]+")
	filteredStr := reg.ReplaceAllString(text, "")
	return filteredStr
}
