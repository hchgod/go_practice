package main

import "fmt"

func maxSubArray(nums []int) int {
	res := nums[0]
	max := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i]
		if res > max {
			max = res
		}
		if res < 0 {
			res = 0
		}
	}
	fmt.Println(max)
	return max
}

func main() {
	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}
