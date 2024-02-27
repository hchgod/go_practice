package main

import (
	"fmt"
)

func violenceSearchInsert(nums []int, target int) int {
	for i :=range nums {
		if nums[i] == target {
			return i
		}else if target>nums[i] && target<nums[i+1]&& i+1 < len(nums) {
			return i+1
		}else {
			break
		}
	}
	return len(nums)
}


func binarySearchInsert(nums []int, target int) int {
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
	return left
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 8
	// result := violenceSearchInsert(nums, target)
	result := binarySearchInsert(nums, target)
	fmt.Println(result)
}

