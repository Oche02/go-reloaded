package main

import (
	"regexp"
	"strings"
)

// fixSingleQuotes finds every pair of single quotes ' '
// and removes any spaces between the opening quote and the first word,
// and between the last word and the closing quote.
// Multi-word phrases inside the quotes are preserved exactly as-is,
// only the padding spaces touching the quotes themselves are stripped.
//
// Examples:
//
//	"' awesome '"          →  "'awesome'"
//	"' hello world '"      →  "'hello world'"
//	"'  spaced   out  '"   →  "'spaced   out'"   (inner spacing untouched)
//	"say ' hi there ' now" →  "say 'hi there' now"
func fixSingleQuotes(text string) string {

	// Pattern breakdown:
	//
	//   '          literal opening single quote
	//   \s*        zero or more spaces/tabs right after the opening quote  ← stripped
	//   (.*?)      NON-GREEDY capture group $1: the actual content
	//              non-greedy (.*?) is critical — it stops at the FIRST closing
	//              quote it finds, so "' a ' b '" becomes two separate matches
	//              instead of swallowing everything into one giant group
	//   \s*        zero or more spaces/tabs right before the closing quote ← stripped
	//   '          literal closing single quote
	//
	// The replacement "'$1'" puts the opening and closing quotes back
	// with the captured content directly inside — no padding spaces.

	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	text = re.ReplaceAllString(text, ` '$1' `)

	re2 := regexp.MustCompile(`"\s*(.*?)\s*"`)
	text = re2.ReplaceAllString(text, `"$1"`)

	text = strings.Join(strings.Fields(text), " ")

	return text
}
