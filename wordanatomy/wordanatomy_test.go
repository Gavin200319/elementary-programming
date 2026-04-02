package main

import (
	"testing"
)

// TestWordAnatomy runs multiple test cases for the WordAnatomy function.
func TestWordAnatomy(t *testing.T) {
	prefixes := []string{"un", "re", "in", "im", "dis", "pre"}
	suffixes := []string{"able", "ible", "ness", "ment", "tion", "er", "ing", "ly"}

	testCases := []struct {
		word     string
		prefixes []string
		suffixes []string
	}{
		{"unbelievable", prefixes, suffixes},
		{"redoing", prefixes, suffixes},
		{"impossible", prefixes, suffixes},
		{"happiness", prefixes, suffixes},
		{"preorder", prefixes, suffixes},
		{"convert", prefixes, suffixes},
		{"dismount", prefixes, suffixes},
		{"unlikely", prefixes, suffixes},
		{"starter", prefixes, suffixes},
	}

	// Optional: generate a random test word
	lastWord := prefixes[0] + "random" + suffixes[0]
	testCases = append(testCases, struct {
		word     string
		prefixes []string
		suffixes []string
	}{lastWord, prefixes, suffixes})

	for _, tc := range testCases {
		t.Run(tc.word, func(t *testing.T) {
			result := WordAnatomy(tc.word, tc.prefixes, tc.suffixes)
			if result == nil {
				t.Errorf("WordAnatomy(%q) returned nil, expected a slice of strings", tc.word)
			} else if len(result) == 0 {
				t.Errorf("WordAnatomy(%q) returned an empty slice", tc.word)
			}
		})
	}
}
