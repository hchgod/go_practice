package main

import "fmt"

func fibonacci(n int) int {
	if n == 0{
		return 0
	}
	a, b, c := 0, 1, 0
	for i := 1; i < n; i++ {
		c = a + b
		a = b
		b = c
	}
	return b
}

func main() {
	fmt.Println(fibonacci(10))
}
