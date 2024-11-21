package main

import (
	//"fmt"
)

// 贪心
// func wiggleMaxLength(nums []int) int {
// 	var res int
// 	if len(nums) <= 1 {
// 		return len(nums)
// 	}
// 	if nums[1] == nums[0] {
// 		res = 1
// 	} else {
// 		res = 2
// 	}
// 	prediff := nums[1] - nums[0]
// 	for i := 2; i < len(nums); i++ {
// 		postdiff := nums[i] - nums[i-1]
// 		if prediff >=0 && postdiff < 0 ||  prediff <=0 && postdiff > 0{
// 			res++
// 			prediff = postdiff
// 		}
// 		fmt.Println("i is ", nums[i], "res is", res)
// 	}
// 	fmt.Println(res)
// 	return res
// }

// 动规
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	if nums[1] == nums[0] {
		dp[1] = 1
	} else {
		dp[1] = 2
	}
	pre := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		post := nums[i] - nums[i-1]
		if pre >= 0 && post < 0 || pre <= 0 && post > 0 {
			dp[i] = dp[i-1] + 1
			pre = post
		}else{
			dp[i] = dp[i-1]
		}
	}

	return dp[len(nums) - 1]
}

func main() {
	wiggleMaxLength([]int{3, 3, 3, 2, 5})
	wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8})
}
