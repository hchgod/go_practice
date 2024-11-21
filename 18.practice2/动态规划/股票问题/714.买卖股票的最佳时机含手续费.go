package main

//可以多次买卖股票  含有手续费  就是每次买卖都需要fee手续费
import (
	"fmt"
)

func maxProfit(prices []int, fee int) int {
	len := len(prices)
	max_value := 0
	dp := make([][]int, len)
	for i := 0; i < len; i++ {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = -fee
	for i := 1; i < len; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i], dp[i-1][2]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][2])
		dp[i][2] = dp[i-1][0] + prices[i] - fee
	}
	for i := 0; i < len; i++ {
		fmt.Println(dp[i])
	}
	for j := 1; j < 3; j++ {
		if dp[len-1][j] > max_value {
			max_value = dp[len-1][j]
		}
	}
	return max_value
}

func main() {
	maxProfit([]int{1, 3, 2, 8, 4, 9}, 2)
}
