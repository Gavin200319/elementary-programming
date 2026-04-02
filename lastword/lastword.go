package main

import (
	"fmt"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
func LastWord(s string) string {
    // Remove trailing spaces
    i := len(s) - 1
    for i >= 0 && s[i] == ' ' {
        i--
    }

    if i < 0 { // empty or all spaces
        return "\n"
    }

    // Find the start of the last word
    end := i
    for i >= 0 && s[i] != ' ' {
        i--
    }

    return s[i+1 : end+1] + "\n"
}

