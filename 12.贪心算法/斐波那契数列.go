package main

import (
	"fmt"
)

func fib(num int) int {
	res := 0
	if num == 0 || num == 1 {
		return num
	}
	res = fib(num-1) + fib(num-2)
	return res
}


func main() {
	a := 4
	fmt.Println(fib(a))
}
