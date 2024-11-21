package main

//可以多次买卖股票  含有冷冻期   就是卖出后第二天不能买
import (
	"fmt"
)

func maxProfit(prices []int) int {
	len := len(prices)
	max_value := 0
	dp := make([][]int, len)
	for i := 0; i < len; i++ {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0
	dp[0][3] = 0
	for i := 1; i < len; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i], dp[i-1][3]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	for i := 0; i < len; i++ {
		fmt.Println(dp[i])
	}
	for j := 0; j < 4; j++ {
		if dp[len-1][j] > max_value {
			max_value = dp[len-1][j]
		}
	}
	return max_value
}

func main() {
	maxProfit([]int{1, 2, 3, 0, 2})
}
