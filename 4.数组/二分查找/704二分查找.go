package main

import (
	"fmt"
)

func binary(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left +(right-left) / 2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else if target == nums[mid] {
			return mid
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := -1
	result := binary(nums, target)
	fmt.Println(result)
}

