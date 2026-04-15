package main

import "strings"

func processText(text string) string {
	words := strings.Fields(text)

	words = hexToDecimal(words)
	words = binToDecimal(words)
	words = toUpper(words)
	words = toLower(words)

	lines := strings.Join(words, " ")

	return lines
}
