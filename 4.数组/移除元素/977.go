package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	k := len(nums) - 1
	result := make([]int, len(nums))
	for left <= right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			result[k] = nums[right] * nums[right]
			right--
		} else if nums[left]*nums[left] >= nums[right]*nums[right] {
			result[k] = nums[left] * nums[left]
			left++
		}
		k--
	}
	return result
}

func main() {
	nums := []int{-4,-1,0,3,10}
	nums_return := sortedSquares(nums)
	fmt.Println(nums_return)
}
