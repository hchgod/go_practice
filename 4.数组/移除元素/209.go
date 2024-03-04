package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	slow, fast, sum, result := 0, 0, 0, len(nums)+1
	temp := 0
	for ; fast < len(nums); fast++ {
		sum += nums[fast]
		for sum >= target {
			temp = fast - slow + 1
			if temp < result {
				result = temp
			}
			sum -= nums[slow]
			slow++
		}
	}
	if result == len(nums)+1 {
		return 0
	} else {
		return result
	}
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	nums_return := minSubArrayLen(target, nums)
	fmt.Println(nums_return)
}
