package main

import (
	"fmt"
)

func isHappy(n int) bool {
	res := make(map[int]bool)
	for n!=1 && res[n]!=true {
		res[n] = true
		n = sum(n)
	}
	return n==1
}

func sum(n int) int {
    sum := 0
    for n > 0 {
        sum += (n % 10) * (n % 10)
        n = n / 10
    }
    return sum
}

func main() {
	n := 19
	res := isHappy(n)
	fmt.Print(res)
}
