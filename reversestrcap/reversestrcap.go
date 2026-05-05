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
	r:=[]rune(s)
	for i:=0; i<len(r);i++{
		r[i]= toLower(r[i])
	}
	for i:=0; i<len(r); i++{
		if isLetter(r[i]){
			if i+1 == len(r)|| !isLetter(r[i+1]){
				r[i]= toUpper(r[i])
			}
		}
	}
	return string (r)
}
