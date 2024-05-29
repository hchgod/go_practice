package main

import "fmt"

func canJump(nums []int) bool {
	cover := nums[0]
	n := len(nums)
	for i := 1; i <= cover; i++ {
		if i+nums[i] > n {
			return true
		} else {
			cover = max(cover, i+nums[i])
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	a := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(a))
}
