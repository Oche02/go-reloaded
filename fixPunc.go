package main

import (
	"regexp"
)

func fixPunctuation(text string) string {
	// remove multiples of spaces and replace it with just a space
	spacesReplace := regexp.MustCompile(` +`)
	text = spacesReplace.ReplaceAllString(text, " ")

	// remove space or spaces before punctuation mark
	beforePunc := regexp.MustCompile(` +([,.?;:'"!#]+)`)
	text = beforePunc.ReplaceAllString(text, ("$1"))

	// remove spaces after punctuation mark when the next character as a letter or number
	afterPunc := regexp.MustCompile(`([.,;:"1{}])`)
	text = afterPunc.ReplaceAllString(text, "$1 $2")

	return text
}
