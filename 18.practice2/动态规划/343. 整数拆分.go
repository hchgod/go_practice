package main

import (
	"fmt"
)

func function(n int) int{
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			fmt.Println("此时此刻i和j的值:",i,j)
			dp[i] = max(max(j*(i-j),dp[i-j]*j),dp[i])
			fmt.Println(dp)
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func main() {
	function(10)
}
