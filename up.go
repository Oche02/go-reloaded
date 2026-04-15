package main

import "strings"

func toUpper(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])

			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
