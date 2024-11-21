package main
//最多只能买卖k次股票
import (
	"fmt"
)

func maxProfit(k int, prices []int) int {
	len := len(prices)
	max_value := 0
	dp := make([][]int,len)
	for i := 0; i < len; i++{
		dp[i] = make([]int, 2*k+1)
	}
	for i := 0; i < len; i++{
		dp[i][0] = 0
	}
	for j := 0; j < k; j++{
		dp[0][2*j+1] = -prices[0]
		dp[0][2*j+2] = 0
	}
	for i := 1; i < len; i++{
		for j := 0; j < k; j++{
			dp[i][2*j+1] = max(dp[i-1][2*j+1],dp[i-1][2*j]-prices[i])
			dp[i][2*j+2] = max(dp[i-1][2*j+2],dp[i-1][2*j+1]+prices[i])
		}
	}
	for j := 0; j <= 2*k; j++{
		if dp[len-1][j] > max_value {
			max_value = dp[len-1][j]
		}
	}
	fmt.Println(dp)
	return max_value
}

func main() {
	maxProfit(2,[]int{2,4,1})
}