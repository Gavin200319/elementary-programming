package student

import "testing"

func TestCountRepeats(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"ABCABC", "ABCABC"},
		{"AAABBC", "A3B2C"},
		{"JjjJohhnnnNn", "Jj2Joh2n3Nn"},
		{"     ", " 5"},
		{" AAABBC ", " A3B2C "},
		{"", ""},
		{"A", "A"},
		{"AA", "A2"},
		{"AAAAAA", "A6"},
		{"123333", "1234"},
		{"aAbBCc", "aAbBCc"},
		{"!!@@##", "!2@2#2"},
		{"   AAA   ", " 3A3 3"},
		{"111ABC", "13ABC"},
		{"aaaBBBccc1111", "a3B3c314"},
	}

	for _, tt := range tests {
		result := CountRepeats(tt.input)
		if result != tt.expected {
			t.Errorf("CountRepeats(%q) = %q, expected %q", tt.input, result, tt.expected)
		}
	}
}
