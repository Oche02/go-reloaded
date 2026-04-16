package main

import (
	"strconv"
	"strings"
)

func toUpper(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])

			words = append(words[:i], words[i+1:]...)
		}

		if words[i] == "(up," && i+1 <= len(words) {
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err == nil {
				for j := 1; j <= n && i-j >= 0; j++ {
					words[i-j] = strings.ToUpper(words[i-j])
				}
			}
			words = append(words[:i], words[i+2:]...)
		}
	}
	return words
}
