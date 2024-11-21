package main
//只能买卖一次股票
import (
	"fmt"
	"math"
)

//双指针
func function(prices []int) {
	low := math.MaxInt32

	res := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < low {
			low = prices[i]
		}
		res  = max((prices[i] - low),res)
	}
	fmt.Println(res)
	return
}

func maxProfit(prices []int) int {
	len := len(prices)
	dp := make([][]int,len)
	for i := 0; i < len; i++{
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len; i++{
		dp[i][0] = max(dp[i-1][0],-prices[i])
		dp[i][1] = max(dp[i-1][1],prices[i]-prices[i-1])
	}
	fmt.Println(dp)
	return dp[len-1][1]
}

func main() {
	function([]int{7,6,4,3,1})
	maxProfit([]int{7,1,5,3,6,4})
}