package main

import (
	"fmt"
)

func moveZeroes(nums []int) []int{
	slow,fast := 0,0
	temp := 0
	for ;fast < len(nums); fast++ {
		if nums[slow] == 0&&nums[fast] != 0{
			temp = nums[fast]
			nums[fast]=nums[slow]
			nums[slow]=temp
			slow++
		}else if nums[slow]!=0{
			slow++
		}
	}
	return nums
}


func main() {
	nums := []int{0,1,0,3,12}
	nums_return := moveZeroes(nums)
	fmt.Println(nums_return)
}