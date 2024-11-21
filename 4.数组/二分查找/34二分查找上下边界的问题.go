package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left +(right-left) / 2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else if target == nums[mid] {
			count:=0
			for i:=0 ; nums[mid]==nums[mid+i]; i++{
				count++
			}
			return []int{mid,mid+count}
		}
	}
	return []int{-1,-1}
}

func main() {
	nums := []int{5,7,7,8,8,10}
	target := 7
	result := searchInsert(nums, target)
	fmt.Println(result)
}

