package main

import (
	"fmt"
)

func total_Sqrt(target int) bool {
	left, right := 0, target
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid > target {
			right = mid - 1
		} else if mid*mid < target {
			left = mid + 1
		} else if mid*mid == target {
			return true
		}
	}
	return false
}

func main() {
	target := 1
	result := total_Sqrt(target)
	fmt.Println(result)
}
