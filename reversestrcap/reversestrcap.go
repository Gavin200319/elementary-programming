package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println()
		return
	}
	for _, v := range args {
		fmt.Println(process(v))
	}
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func strCap(s string) string {
	r := []rune(s)
	inWord := false

	for i := 0; i < len(r); i++ {
		if isLetter(r[i]) {
			if !inWord {
				r[i] = toUpper(r[i])
				inWord = true
			} else {
				r[i] = toLower(r[i])

			}
		} else {
			inWord = false
		}
	}
	return string(r)
}

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
func toUpper(c rune) rune {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}
	return c
}
func toLower(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
func process(s string) string {
	return reverse(strCap(s))
}
