package main

import "fmt"

func fibonacci(n int) int {
	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	fmt.Println(b)
	return b

}

func main() {
	fibonacci(5)
}
