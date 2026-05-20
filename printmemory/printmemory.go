package main

import "github.com/01-edu/z01"

func printmemory(arr [10]byte){
	//loop in chunks of 4
	for i:=0; i<len(arr); i+=4{
		
		//print hex value
		for j:= 0; j<len(arr); j++{
			if i+j <len(arr){
				z01.PrintRune(' ')
				printHexByte(arr[i+j])
			}else{
				//space separator
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				z01.PrintRune(' ')
			}

		}
		//print ASCII characters
		for j:=0; j<4;j++{
			if i+j<len(arr){
				c:=arr[i+j]
				if c>32 && c<=126{
					z01.PrintRune(rune(c))
				}else{
					z01.PrintRune('.')
				}
			}
		}
		z01.PrintRune('\n')
	}

}

func printHexByte(b byte){
		hex:= "0123456789abcdef"
		z01.PrintRune(rune(hex[b/16]))
		z01.PrintRune(rune(hex[b%16]))
	}

func main() {
	// Test case 1
	arr1 := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}
	printmemory(arr1)

	// Test case 2
	arr2 := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	printmemory(arr2)

	// Test case 3
	arr3 := [10]byte{}
	printmemory(arr3)
}