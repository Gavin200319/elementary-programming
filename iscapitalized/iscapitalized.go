package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

func IsCapitalized(s string) bool {
	if len(s)==0{
		return false
	}
	startOfWord := true

	for _, r := range s{
		if r == ' '{
			startOfWord=true
			continue
		}
		if startOfWord{
			if r >= 'a' && r <= 'z'{
				return false
			}
			startOfWord = false
		}
	}
	return true
}