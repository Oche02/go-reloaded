package main

import (
	"regexp"
	"strings"
)

func fixPunctuation(text string) string {
	spaceRe := regexp.MustCompile(` +`)
	text = spaceRe.ReplaceAllString(text, " ")

	beforePunct := regexp.MustCompile(` +([.,!?:;]+)`)
	text = beforePunct.ReplaceAllString(text, "$1")

	afterPunct := regexp.MustCompile(`([.,!?:;]+)([^\s.,!?:;'"])`)
	text = afterPunct.ReplaceAllString(text, "$1 $2")

	// Step 4: Trim any leading/trailing spaces per line
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return strings.Join(lines, "\n")
}
