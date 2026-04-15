package main

import "strings"

func processText(text string) string {
	words := strings.Fields(text)

	words = hexToDecimal(words)

	lines := strings.Join(words, " ")

	return lines
}
