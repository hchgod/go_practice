package main

import (
	"fmt"
)

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	hashmap := make(map[int]int)
	count := 0
	for _,val1 := range nums1 {
		for _,val2 := range nums2 {
			hashmap[val1+val2]++
		}
	}
	for _,val3 := range nums3 {
		for _,val4 := range nums4 {
			count = count + hashmap[-(val3+val4)]
		}
	}
	return count
}

func main() {
	n := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(n,target)
	fmt.Print(res)
}
