package main

import (
	"regexp"
	"strings"
)

func fixQuotation(text string) string {
	quotBeforeText := regexp.MustCompile(`'["()^"]'`)
	text = quotBeforeText.ReplaceAllStringFunc(text, func(match string) string {
		inner := match[1 : len(match)-1]
		inner = strings.TrimSpace(inner)
		return "'" + inner + "'"
	})
	return text
}
