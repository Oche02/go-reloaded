package main

import "strings"

var vowelsAndH = "aeiouhAEIOUH"

func fixArticlesA(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" || words[i] == "A" {
			nextWord := words[i+1]
			if len(nextWord) > 0 && strings.ContainsRune(vowelsAndH, rune(nextWord[0])) {
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

var consonant = "bcdfgjhklmnpqrstvwxyzBCDFGJHKLMNPQRSTVWXYZ"

func fixArticlesB(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "an" || words[i] == "An" {
			nextWord := words[i+1]
			if len(nextWord) > 0 && strings.ContainsRune(consonant, rune(nextWord[0])) {
				if words[i] == "An" {
					words[i] = "A"
				} else {
					words[i] = "a"
				}
			}
		}
	}
	return words

}
