package main

import "fmt"

func bag01(value, weight []int, bagWeight int) int {
	dp := make([][]int, len(value))    //二维数组的初始化
	for i := range weight {
		dp[i] = make([]int, bagWeight+1)    //注意这个地方的bagweight+1
	}
	for i := 0; i <= bagWeight; i++ {
		if i >= weight[0] {
			dp[0][i] = value[0]
		}
	}
	// fmt.Println(dp)
	for i := 1; i < len(value); i++ {
		for j := 1; j <= bagWeight; j++ {
			if weight[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	fmt.Println(dp)
	return dp[len(value)-1][bagWeight]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	value := []int{15, 20, 30}
	weight := []int{1, 3, 4}
	bag := 4
	fmt.Println(bag01(value, weight, bag))
}
