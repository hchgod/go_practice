package main
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的最短子数组，并返回其长度。
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
