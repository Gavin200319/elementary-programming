package main

import (
	"os"
	"github.com/01-edu/z01"
)


func main(){
if len(os.Args) != 2{
	return
}

s:= os.Args[1]
firstword :=true
i:=0

for i<len(s) {

	for i< len(s) && (s[i]== ' ' || s[i]=='\t'){
		i++
	}
	if i>= len(s){
		break
	}
	// print separator before every word except the first
	if !firstword {
		z01.PrintRune(' ')
		z01.PrintRune(' ')
		z01.PrintRune(' ')
	}

	firstword= false

	//print the word
	for i <len(s) && s[i] != ' ' && s[i] != '\t'{
		z01.PrintRune(rune(s[i]))
		i++
	}
}
	if !firstword{
		z01.PrintRune('\n')
	}
}

