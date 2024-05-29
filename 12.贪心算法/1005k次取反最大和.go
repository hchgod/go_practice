package main

import (
	"fmt"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	if nums[0] >= 0 {
		nums[0] = -nums[0]
		return sum(nums)
	}
	for i := 0; k>0; k-- {
		if nums[i] >= 0 {
			if i > 0 && nums[i] > nums[i-1]{
				nums[i-1] = -nums[i-1]
				continue
			}
			nums[i] = -nums[i]
			continue
		}
		nums[i] = -nums[i]
		i++
		if i == len(nums){
			sort.Ints(nums)
			i = 0
		}
	}
	return sum(nums)
}

func sum(slice []int) int {
	total := 0
	for _, value := range slice {
		total += value
	}
	return total
}

func main() {
	a := []int{3, -1, 0, 2}
	a1 := []int{-2, 9, 9, 8, 4}
	a2 := []int{-4,-2,-3}
	fmt.Println(largestSumAfterKNegations(a, 3))
	fmt.Println(largestSumAfterKNegations(a1, 5))
	fmt.Println(largestSumAfterKNegations(a2, 4))
}
