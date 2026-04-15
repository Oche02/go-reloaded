package main

import "strings"

func toLower(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1])

			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
