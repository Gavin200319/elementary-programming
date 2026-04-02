package main

import (
	"testing"
)

func TestLastWord(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"this        ...       is sparta, then again, maybe    not", "not\n"},
		{" lorem,ipsum ", "lorem,ipsum\n"},
		{" ", "\n"},
		{"hello", "hello\n"},
		{"  multiple   spaces  here  ", "here\n"},
		{"trailing spaces   ", "spaces\n"},
		{"", "\n"},
	}

	for _, tt := range tests {
		result := LastWord(tt.input)
		if result != tt.expected {
			t.Errorf("LastWord(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
