package main

import "fmt"

func function(n int) bool {
	nums := []int{1,5,11,5}
	if Sum(nums)%2 != 0{
		return false
	}
	dp := make([]int,Sum(nums)/2+1)
	for i := 0; i < 4; i++{
		for j := Sum(nums)/2; j > 0; j--{
			if j-nums[i] >= 0{
				dp[j] = max(dp[j],dp[j-nums[i]]+nums[i])
			}
		}
		fmt.Println(dp)
	}
	return dp[len(dp)-1] == Sum(nums)/2
}

func Sum(arr []int) int {
	sum := 0
	for i := range arr{
		sum = sum + arr[i]
	}
	return sum
}

func main() {
	fmt.Println(function(5))
}