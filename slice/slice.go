package main

import "fmt"

func Slice(a []string, nbrs ...int) []string {
	n:= len(a)
	if len(nbrs)==0{
		 return nil
	}
	//set start and end
	start := nbrs[0]
	end :=n
	//if second number exists
	if len(nbrs)>1{
		end = nbrs[1]
	}
	//handle negative indices
	if start<0{
		start = n+ start
	}
	if end<0{
		end = n+ end
	}
	//fix boundaries
	if start <0{
		start =0
	}
	if end>n{
		end = n
	}
	//invalid range check
	if start >= end{
		return nil
	}
	return a[start:end]
}

func main(){
    a := []string{"coding", "algorithm", "ascii", "package", "golang"}
    fmt.Printf("%#v\n", Slice(a, 1))
    fmt.Printf("%#v\n", Slice(a, 2, 4))
    fmt.Printf("%#v\n", Slice(a, -3))
    fmt.Printf("%#v\n", Slice(a, -2, -1))
    fmt.Printf("%#v\n", Slice(a, 2, 0))
}