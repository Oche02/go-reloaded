package main

import "strings"

func processText(text string) string {
	words := strings.Fields(text)

	words = hexToDecimal(words)
	words = binToDecimal(words)
	words = toUpper(words)
	words = toLower(words)
	words = capitalize(words)
	words = fixArticles(words)

	lines := strings.Join(words, " ")
	lines = fixSingleQuotes(lines)
	lines = fixPunctuation(lines)

	return lines
}
