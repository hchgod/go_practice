package main
//多次买卖股票
import (
	"fmt"
)

//双指针法
func function(prices []int) int{
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] - prices[i] > 0 {
			res = res + prices[i+1] - prices[i]
		}
	}
	fmt.Println(res)
	return res
}

func main() {
	// function([]int{7,6,4,3,1})
	maxProfit1([]int{1,7,4,2})
}

// //2列   持有和不持有
// func maxProfit(prices []int) int {
// 	length := len(prices)
//     dp := make([][]int, length)
// 	for i := 0; i < length; i++{
// 		dp[i] = make([]int, 2)
// 	}
// 	dp[0][0] = -prices[0]
// 	dp[0][1] = 0
// 	for i := 1; i < length; i++{
// 		dp[i][0] = max(dp[i-1][0],dp[i-1][1]-prices[i])
// 		dp[i][1] = max(dp[i-1][1],dp[i-1][0]+prices[i])
// 	}
// 	fmt.Println(dp[length - 1][1])
//     return dp[length - 1][1]
// }

//4列   保持持有，今日买入，保持不持有， 今日卖出
func maxProfit1(prices []int) int {
	length := len(prices)
	max_value := 0
    dp := make([][]int, length)
	for i := 0; i < length; i++{
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = -prices[0]
	for i := 1; i < length; i++{
		dp[i][0] = max(dp[i-1][2]-prices[i], dp[i-1][3]-prices[i])
		dp[i][1] = max(dp[i-1][1],dp[i-1][0])
		dp[i][2] = max(dp[i-1][1]+prices[i], dp[i-1][0]+prices[i])
		dp[i][3] = max(dp[i-1][2],dp[i-1][3])
	}
	for j := 2; j < 4; j++ {
		if dp[length-1][j] > max_value {
			max_value = dp[length-1][j]
		}
	}

	fmt.Println(dp)
    return dp[length - 1][1]
}
