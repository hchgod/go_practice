package main

import (
	"fmt"
)

func x_Sqrt(target int) int {
	left, right := 0, target/2
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid > target {
			right = mid - 1
		} else if mid*mid < target {
			left = mid + 1
		} else if mid*mid == target {
			return mid
		}
	}
	return right
}

func main() {
	target := 16
	result := x_Sqrt(target)
	fmt.Println(result)
}
