package main

import "fmt"

func integerBreak(n int) int {
	res := make([]int, n+1)
	res[1] = 1
	res[2] = 1
	for i := 3; i < n+1; i++ {
		for j := 1; j < i; j++ {
			res[i] = max(res[i],max(j*res[i-j],(i-j)*j))
		}
	}
	return res[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(integerBreak(10))
}
