package main

import "strconv"

func binToDecimal(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" && i > 0 {
			num, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(num, 10)

				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	return words
}
