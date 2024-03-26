package main

import "fmt"

func main() {
	res := func(a int,b int)bool{
		if a>b{
			return true
		}else{
			return false
		}
	}
	fmt.Print(res)

}