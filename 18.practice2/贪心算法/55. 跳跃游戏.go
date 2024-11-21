package main

import (
	"fmt"
)

func Jump(nums []int) bool {
	len1 := len(nums)
	if len1 <= 1 {
		return false
	}
	dp := make([]int,len1)
	dp[0] = 0
	cover := 0
	for i := 0; i <= cover; i++ {
		cover = max(i + nums[i], cover)
		fmt.Println("i is",nums[i],"cover is ",cover)
		if cover >= len1-1 {
			return true
		}
	}
	return false
}

func main() {
	Jump([]int{2,3,1,1,4})
	Jump([]int{0,2,3})
}