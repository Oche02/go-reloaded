package main

import (
	"regexp"
	"strings"
)

// FixPunctuation is like a professional editor with a red pen.
// Rules:
//   1. Punctuation (. , ! ? : ;) must hug the word BEFORE it — no gap.
//   2. There must be a space AFTER punctuation (before the next word).
//   3. Groups like "..." or "!?" stay together and follow the same rules.
//
// Think of it like parking rules: punctuation always parks RIGHT NEXT
// to the previous word (no empty space), with room after it for the
// next word to park.
//
// "Hello , world !" → "Hello, world!"
// "I was thinking ... You were right" → "I was thinking... You were right"

func fixPunctuation(text string) string {
	// Step 1: Collapse multiple spaces into one
	spaceRe := regexp.MustCompile(` +`)
	text = spaceRe.ReplaceAllString(text, " ")

	// Step 2: Remove spaces BEFORE punctuation groups
	// e.g. "Hello ," → "Hello,"
	// e.g. "Wait ..." → "Wait..."
	// Matches: optional spaces, then one or more of . , ! ? : ;
	beforePunct := regexp.MustCompile(` +([.,!?:;]+)`)
	text = beforePunct.ReplaceAllString(text, "$1")

	// Step 3: Ensure exactly one space AFTER punctuation groups
	// (but only if the next character is a letter/digit, not end of string)
	// e.g. "Hello,world" → "Hello, world"
	// e.g. "Wait...You" → "Wait... You"
	afterPunct := regexp.MustCompile(`([.,!?:;]+)([^\s.,!?:;'"])`)
	text = afterPunct.ReplaceAllString(text, "$1 $2")

	// Step 4: Trim any leading/trailing spaces per line
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return strings.Join(lines, "\n")
}
