package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main(){
	if len(os.Args) !=3{
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]

	if s1 ==""{
		z01.PrintRune('1')
		z01.PrintRune('\n')
	}

	i:=0 //pointers for s1

	for j:=0; j<len(s2) && i< len(s1); j++{
		if s1[i]==s2[j]{
			i++
		}
	}

	if i==len(s1){
		z01.PrintRune('1')
	}else{
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')

}

