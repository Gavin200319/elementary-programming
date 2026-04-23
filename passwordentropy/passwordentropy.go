package student

import "math"

// Entropy calculates the entropy of a password.
// Formula: entropy = length * log2(charset size)
func Entropy(password string) float64 {
	if len(password) == 0 {
		return 0
	}

	charsetSize := 0

	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSymbol := false

	for _, c := range password {
		switch {
		case c >= 'a' && c <= 'z':
			hasLower = true
		case c >= 'A' && c <= 'Z':
			hasUpper = true
		case c >= '0' && c <= '9':
			hasDigit = true
		default:
			hasSymbol = true
		}
	}

	if hasLower {
		charsetSize += 26
	}
	if hasUpper {
		charsetSize += 26
	}
	if hasDigit {
		charsetSize += 10
	}
	if hasSymbol {
		// Common printable symbols (approximation)
		charsetSize += 32
	}

	if charsetSize == 0 {
		return 0
	}

	entropy := float64(len(password)) * math.Log2(float64(charsetSize))
	return entropy
}
