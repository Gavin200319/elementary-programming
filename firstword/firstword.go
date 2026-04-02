
package main

import (
	"fmt"
)

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}
func FirstWord(s string) string {
	i := 0

	// skip leading spaces
	for i < len(s) && s[i] == ' ' {
		i++
	}

	start := i

	// find end of first word
	for i < len(s) && s[i] != ' ' {
		i++
	}

	if start == len(s) {
		return "\n"
	}

	return s[start:i] + "\n"
}
