package main
//最长递增的连续递增序列  dp数组表示以i为结尾的递增子数组长度
import "fmt"

//动态规划的解法
func findLengthOfLCIS(nums []int) int {
	maxvalue := 0
	len := len(nums)
	dp := make([]int, len)
	for i := 0; i < len; i++ {
		dp[i] = 1
	}
	for i := 1; i < len; i++ {
		if nums[i-1] < nums[i] {
			// dp[i] = max(dp[i-1] +1,dp[i])
			dp[i] = dp[i-1] +1
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

//正常循环的解法
func findLengthOfLCIS1(nums []int) int {
	maxl := 1
	temp := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			temp++
			maxl = max(maxl,temp)
		}else{
			temp = 1
		}
	}
	return maxl
}


func main() {
	findLengthOfLCIS([]int{1,3,5,4,7})
}
