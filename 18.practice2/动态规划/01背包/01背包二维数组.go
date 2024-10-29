package main

import "fmt"

func function(n int) {
	bag := 4
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int,bag+1)
	}
	for j := 0; j < bag+1; j++ {
		if weight[0] <= j{
			dp[0][j] = value[0]
		}
	}
	for i := 1; i < len(value); i++ {
		for j := 1; j < bag+1; j++ {
			if j >= weight[i]{
				dp[i][j] = max(dp[i-1][j],dp[i][j-weight[i]]+value[i])
			}else{
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	fmt.Println(dp)
	return
}

func main() {
	function(5)
}
