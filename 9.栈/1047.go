package main

import (
	"fmt"
)

//"fmt"

func removeDuplicates(s string) string {
	if s == "" {
		return s
	}
	str := []byte(s)
	stack := []byte{}
	for _, v := range str {
		if len(stack)>0&&stack[len(stack)-1] == v{
			stack = stack[:len(stack)-1]
		}else{
			stack = append(stack,byte(v))
		}
	}
	return string(stack)
}

func main() {
	s := "abbaca"
	res := removeDuplicates(s)
	fmt.Print(res)
}
