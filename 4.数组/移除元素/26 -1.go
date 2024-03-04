package main

import(
	"fmt"
)

func removeDuplicate(nums []int) int {
	slow := 0
	count := 1
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			count++
			slow++
			nums[slow] = nums[fast]						
		}
	}
	nums = nums[:count]
	fmt.Println(nums)
	return len(nums[0:count])
}

func main() {
	nums := []int{}
	nums_len := removeDuplicate(nums)
	fmt.Println(nums_len)
}