package main

import "fmt"

func fifthandskip(str string) string{
	//handle empty string
	if str== "" {
		return "\n"
	}
	// initialize variables
	result := ""
	count :=0

	//loop through the string
	for i:=0; i<len(str); i++{
		if str[i]==' '{
			continue
		}

		count++
		//take first 5 characters
		if count <=5{
			result += string(str[i])
		}
		//skip the 6th character
		if count ==6{
			count =0
			result += " "
		}		

	}

	//count total valid characters
	cleanCount :=0
	for i:=0; i<len(str); i++{
		if str[i] !=' '{
			cleanCount++
		}
	}
	//handle short inputs
	if cleanCount<5{
		return "Invalid Input\n"
	}
	//remove trailing space
	if len(result)>0 && result[len(result)-1]== ' '{
		result = result[:len(result)-1]
	}
return result +"\n"
}

func main() {
	fmt.Print(fifthandskip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(fifthandskip("This is a short sentence"))
	fmt.Print(fifthandskip("1234"))
}