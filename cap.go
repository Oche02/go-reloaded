package main

import (
	"strconv"
	"strings"
)

func capitalize(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" && i > 0 {
			words[i-1] = capFirst(words[i-1])
			words = append(words[:i], words[i+1:]...)
			continue
		}

		if words[i] == "(cap," && i+1 < len(words) {
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err == nil {
				for j := 1; j <= n && i-j >= 0; j++ {
					words[i-j] = capFirst(words[i-j])
				}
			}
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}

func capFirst(words string) string {
	if len(words) == 0 {
		return words
	}
	return strings.ToUpper(string(words[0])) + words[1:]
}
