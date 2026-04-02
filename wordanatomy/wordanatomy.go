package main

import (
	"fmt"
)
// WordAnatomy splits a word into prefix, root, and suffix
func WordAnatomy(word string, prefixes, suffixes []string) []string {
	prefix, suffix := "", ""

	// Find longest matching prefix
	for _, p := range prefixes {
		if len(word) >= len(p) && word[:len(p)] == p {//&& len(p) > len(prefix) {
			prefix = p
		}
	}

	// Find longest matching suffix
	for _, s := range suffixes {
		if len(word) >= len(s) && word[len(word)-len(s):] == s {//&& len(s) > len(suffix) {
			suffix = s
		}
	}

	// Extract root
	root := word[len(prefix) : len(word)-len(suffix)]

	// Collect non-empty parts
	parts := []string{}
	for _, part := range []string{prefix, root, suffix} {
		if part != "" {
			parts = append(parts, part)
		}
	}

	return parts
}

func main() {
	prefixes := []string{"un", "re", "in", "im", "dis", "pre"}
	suffixes := []string{"able", "ible", "ness", "ment", "tion", "er", "ing", "ly"}

	words := []string{
		"unbelievable",
		"redoing",
		"impossible",
		"happiness",
		"preorder",
		"convert",
		"dismount",
		"unlikely",
		"starter",
	}

	for _, word := range words {
		parts := WordAnatomy(word, prefixes, suffixes)
		fmt.Printf("%s -> %v\n", word, parts)
	}
}
