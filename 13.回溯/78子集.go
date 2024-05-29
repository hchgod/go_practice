package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})
	temp := []int{}
	var backtracing func(start int)
	backtracing = func(start int) {
		for i := start; i < len(nums); i++ {
			temp = append(temp, nums[i])
			copylist := make([]int, len(temp))
			copy(copylist, temp)
			res = append(res, copylist)
			backtracing(i+1)
			temp = temp[:len(temp)-1]
		}
		return
	}
	backtracing(0)
	return res
}

func main() {
	a := []int{1, 2, 3}
	res := subsets(a)
	fmt.Println(res)
}
