package student

import "strconv"

func CountRepeats(s string) string {
	if len(s) == 0 {
		return ""
	}

	result := ""
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			result += string(s[i-1])
			if count > 1 {
				result += strconv.Itoa(count)
			}
			count = 1
		}
	}

	// handle last character/group
	result += string(s[len(s)-1])
	if count > 1 {
		result += strconv.Itoa(count)
	}

	return result
}

