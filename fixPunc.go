package main

import (
	"regexp"
	"strings"
)

// fixPunctuation ensures that punctuation marks . , ! ? : ;
// including GROUPS/RUNS of them (e.g. "...", "!?", "!!")
// are treated as a single unit:
//   - No space before the group
//   - No internal spaces within the group
//   - Exactly one space after the group
//
// Operates on the final joined string, not on a token slice.
func fixPunctuation(text string) string {

	// --- Pass 1: Collapse internal spaces WITHIN a punctuation group ---
	// Problem:  "hello . . ." or "wait ! ?" has spaces splitting what
	//           should be one unit.
	// Pattern:  a punctuation char, then whitespace, then another punctuation char
	//           repeated with + so it handles arbitrarily long runs like ". . . ."
	// Replace:  just the two punctuation chars with no space between them,
	//           the \s+ in the middle is dropped entirely.
	// Example:  "hello . . ."  →  "hello..."
	// Example:  "wait ! ?"     →  "wait!?"
	internalSpaces := regexp.MustCompile(`([.,!?:;])\s+([.,!?:;])`)
	for internalSpaces.MatchString(text) {
		text = internalSpaces.ReplaceAllString(text, "$1$2")
	}

	// --- Pass 2: Remove any whitespace BEFORE a punctuation group ---
	// Now that multi-char groups are collapsed into single tokens (from Pass 1),
	// we can treat the entire run [.,!?:;]+ as one unit.
	// Pattern:  one or more whitespace chars, followed by one or more punctuation chars
	// Replace:  just the punctuation group ($1), the leading whitespace is dropped.
	// Example:  "hello ..."   →  "hello..."
	// Example:  "done !?"     →  "done!?"
	spaceBefore := regexp.MustCompile(`\s+([.,!?:;]+)`)
	text = spaceBefore.ReplaceAllString(text, "$1")

	// --- Pass 3: Ensure exactly ONE space AFTER a punctuation group ---
	// Pattern:  one or more punctuation chars, followed by zero or more spaces,
	//           followed by a non-whitespace character.
	// The \s* (zero or more) handles both missing space AND excess spaces.
	// The \S anchor at the end prevents adding a trailing space at end-of-string.
	// Replace:  punctuation group + single space + the following character.
	// Example:  "hello...world"   →  "hello... world"
	// Example:  "done!!  really"  →  "done!!  really"  →  "done!! really"
	// Example:  "hi!?how"         →  "hi!? how"
	spaceAfter := regexp.MustCompile(`([.,!?:;]+)\s*(\S)`)
	text = spaceAfter.ReplaceAllString(text, "$1 $2")

	// --- Pass 4: Final trim ---
	// Clean up any accidental leading or trailing whitespace on the whole string.
	text = strings.TrimSpace(text)

	return text
}
