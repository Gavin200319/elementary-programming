package main

import (
	"fmt"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {
	runes := []rune(s)
	count := 1
	result := ""
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			count++

		}else{
		result += fmt.Sprintf("%d%c", count, runes[i-1])
		count =1
		}
	}

	return result
}
