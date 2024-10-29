package main

import "fmt"

func function() int {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	n := len(cost)+1
	dp := make([]int, n)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < n; i++ {
		dp[i] = min(cost[i-1]+dp[i-1], cost[i-2]+dp[i-2])
	}
	fmt.Println(dp)
	return dp[n-1]
}

func main() {
	fmt.Println(function())
}
