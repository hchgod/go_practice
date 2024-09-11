package main

import "fmt"

func canPartition(nums []int) bool {
	target := sum(nums)/2
	if sum(nums)%2 != 0{
		return false
	}
	dp := make([]int, target+1) //先初始化一下
	for j := 1; j < len(dp); j++ {
		if j >= nums[0] {
			dp[j] = max(dp[j], nums[0])
		}
	}
	fmt.Println(dp)

	for i := 1; i < len(nums); i++ {
		for j := len(dp)-1; j > 0; j-- {     //这个一定要倒序
			if j >= nums[i] {
				dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
			}
		}
		fmt.Println(dp)
	}
	return dp[target] == target
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	// value := []int{2,3 ,1 ,5 ,4 ,3}
	// weight := []int{2,2 ,3 ,1 ,5 ,2}
	// bag := 1
	num := []int{1,2,5}
	fmt.Println(canPartition(num))
}
