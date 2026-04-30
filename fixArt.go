package main

import (
	"strings"
)

func fixArticles(word []string) []string {

	for i := 0; i < len(word)-1; i++ {
		new := strings.TrimLeft(strings.Join(word[i+1:], ""), "\"',.;:?!")
		if len(new) == 0 {
			continue
		}
		vow := strings.ContainsAny("aeiouhAEIOUH", string(new[0]))
		if word[i] == "a" && vow {
			word[i] = "an"
		} else if word[i] == "A" && vow {
			word[i] = "An"
		} else if word[i] == "an" && !vow {
			word[i] = "a"
		} else if word[i] == "An" && !vow {
			word[i] = "A"
		}

	}
	return word
}
