package helper

import "strings"

var replacer = strings.NewReplacer("/", "_")

func TrimSpecial(source string) string {
	return replacer.Replace(source)
}
