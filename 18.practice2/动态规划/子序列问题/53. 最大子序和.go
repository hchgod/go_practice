package main

import "fmt"

//最大子序和

// 贪心
func maxSubArray(nums []int) int {
    maxl := nums[0]
    sum := nums[0]

    for i := 1; i < len(nums); i++{
		if sum + nums[i] >= 0{
			sum = sum + nums[i]
		} else{
			sum = nums[i]
		}
		fmt.Println("i的值",nums[i], "sum的值",sum)
		maxl = max(sum, maxl)
    }
	fmt.Println(maxl)
    return maxl
}

//dp
// func maxSubArray(nums []int) int {
// 	dp := make([]int, len(nums))
// 	dp[0] = nums[0]
// 	maxl := nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		dp[i] = max(nums[i], dp[i-1]+nums[i])
// 		maxl = max(dp[i], maxl)
// 	}
// 	fmt.Println(dp)
// 	return maxl
// }

func main() {
	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	// maxSubArray([]int{1,-1,1})
	// maxSubArray([]int{-1,-2})
}
