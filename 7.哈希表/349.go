package main

import (
	"fmt"
	"strconv"
)

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	set := make(map[string]int)
	for _, i := range nums1 {
		if set[strconv.Itoa(i)] == 0 {
			set[strconv.Itoa(i)]++
		}
	}
	for _, i := range nums2 {
		if set[strconv.Itoa(i)] != 0 {
			res = append(res, i)
			set[strconv.Itoa(i)]--
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	res := intersection(nums1, nums2)
	fmt.Print(res)
}
