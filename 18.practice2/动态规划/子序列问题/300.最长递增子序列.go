package main
//最长递增的子序列的长度   dp数组表示以i为结尾的递增子数组长度
import "fmt"

func lengthOfLIS(nums []int) int {
	maxvalue := 0
	len := len(nums)
	dp := make([]int, len)
	for i := 0; i < len; i++ {
		dp[i] = 1
	}
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j] +1,dp[i])
			}
		}
		fmt.Println(dp)
	}
	for i := 0; i < len; i++ {
		if dp[i] > maxvalue {
			maxvalue = dp[i]
		}
	}
	fmt.Println(dp)
	return maxvalue
}
func main() {
	lengthOfLIS([]int{0,1,0,3,2,3})
}
