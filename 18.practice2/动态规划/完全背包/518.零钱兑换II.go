package main

import (
	"fmt"
	"sort"
)

func function(n int) int {
	amount := 5
	coins := []int{1, 2, 5}
	sort.Ints(coins)
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < 3; i++ {
		for j := coins[i]; j <= amount; j++ {
			if j >= coins[i] {
				dp[j] = dp[j]+dp[j-coins[i]]
			}
		}
		fmt.Println(dp)
	}
	return dp[amount]
}

func main() {
	function(5)
}
