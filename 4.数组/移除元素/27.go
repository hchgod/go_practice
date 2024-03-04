package main

import(
	"fmt"
)

func removeElement(nums []int, val int) []int {
	slow := 0
	count := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			count ++
			nums[slow] = nums[fast]
			slow++			
		}
	}
	return nums[0:count]
}

func main() {
	nums := []int{}
	value := 3
	nums = removeElement(nums,value)
	fmt.Println(nums)
}