package main

import (
	"fmt"
	"math"
)

func function(n int) {
	dp := make([]int, n+1)
	for j := 1; j <= n; j++{
		dp[j] = math.MaxInt32
	}
	for i := 1; i <= n ; i++ {
		if !isPerfectSquare(i) {
			continue
		}
		for j := 1; j <= n; j++ {
			if j>=i{
				dp[j] = min(dp[j],dp[j-i]+1)
			}
			fmt.Println(dp)
		}
	}
	fmt.Println(dp)
	return
}

func main() {
	function(12)
}

func isPerfectSquare(x int) bool {
	if x < 0 {
		return false
	}
	sqrt := math.Sqrt(float64(x))
	return sqrt == float64(int(sqrt))
}
