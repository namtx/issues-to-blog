package string_helper

import "strings"

var replacer = strings.NewReplacer("/", "_", " ", "_")

func TrimSpecial(source string) string {
	return replacer.Replace(source)
}
