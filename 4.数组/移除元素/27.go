package main

import(
	"fmt"
)

func removeElement(nums []int, val int) []int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	fmt.Println(nums)
	return nums[0:slow]
}

func main() {
	nums := []int{1,2,3,3,4,5}
	value := 3
	nums = removeElement(nums,value)
	fmt.Println(nums)
}