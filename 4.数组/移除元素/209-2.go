package main

import (
	"fmt"
)

func minSubArrayLen2(target int, nums []int) int {
	slow, fast, sum, result := 0, 0, 0, len(nums)+1
	result_left, result_right := 0,0
	temp := 0
	for ; fast < len(nums); fast++ {
		sum += nums[fast]
		for sum >= target {
			temp = fast - slow + 1
			if temp < result {
				result = temp
				result_left = slow
				result_right = fast
			}
			sum -= nums[slow]
			slow++
		}
	}
	if result == len(nums)+1 {
		return 0
	} else {
		fmt.Println(nums[result_left:result_right+1])
		return result
	}
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 4
	nums_return := minSubArrayLen2(target, nums)
	fmt.Println(nums_return)
}
