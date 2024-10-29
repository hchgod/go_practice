package main

import "fmt"

func function(n int) {
	bag := 4
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	dp := make([]int, bag+1)
	for i := 0; i < 3; i++ {
		for j := bag; j >= 0; j-- {
			if weight[i] <= j{
				dp[j] = max(dp[j],dp[j-weight[i]]+value[i])
			}
		}
		fmt.Println(dp)
	}
	return
}

func main() {
	function(5)
}
