package main

import "strings"

func processText(text string) string {
	words := strings.Fields(text)

	words = hexToDecimal(words)
	words = binToDecimal(words)
	words = toUpper(words)
	words = toLower(words)
	words = capitalize(words)
	words = fixArticlesA(words)
	words = fixArticlesB(words)

	lines := strings.Join(words, " ")

	lines = fixPunctuation(lines)
	lines = fixSingleQuotes(lines)

	return lines
}
