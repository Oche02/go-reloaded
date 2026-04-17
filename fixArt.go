package main

import "strings"

func fixArticle(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "a" || words[i] == "A" {
			nextWord := words[i+1]
			if len(nextWord) > 0 && strings.ContainsRune("aeiouhAEIOUH", rune(nextWord[0])) {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}
	return words
}
