package main

import (
	"strconv"
	"strings"
)

func toLower(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1])

			words = append(words[:i], words[i+1:]...)
			continue
		}

		if words[i] == "(low," && i+1 < len(words) {
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err == nil {
				for j := 1; j <= n && i-j >= 0; j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}
			}
			words = append(words[:i], words[i+2:]...)
		}
	}
	return words
}
