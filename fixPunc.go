package main

import (
	"regexp"
)

func fixPunctuation(text string) string {
	spacesReplace := regexp.MustCompile(` +`)
	text = spacesReplace.ReplaceAllString(text, " ")

	return text
}
