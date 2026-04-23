package student

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	const epsilon = 0.0001
	return math.Abs(a-b) < epsilon
}

func TestEntropy(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"", 0},
		{"abc", 3 * math.Log2(26)},
		{"ABC", 3 * math.Log2(26)},
		{"123", 3 * math.Log2(10)},
		{"!@#", 3 * math.Log2(32)},
		{"aA1!", 4 * math.Log2(26+26+10+32)},
		{"password", 8 * math.Log2(26)},
		{"PASSWORD123!", 13 * math.Log2(26+10+32)},
		{"0000", 4 * math.Log2(10)},
	}

	for _, tt := range tests {
		result := Entropy(tt.input)
		if !almostEqual(result, tt.expected) {
			t.Errorf("Entropy(%q) = %f; expected %f", tt.input, result, tt.expected)
		}
	}
}
