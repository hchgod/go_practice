package main

import (
	"fmt"
)

func combinationSum3(n int, k int) [][]int {
	res := [][]int{}
	list := []int{}
	var backtracing func(n int, k int, start int)
	backtracing = func(n int, k int, start int) {
		if len(list) == k {
			if sum(list) == n {
				//fmt.Println(list)
				temp := make([]int, len(list))
				copy(temp, list)
				res = append(res, temp)
			}
			return
		}
		for i := start; i < 10; i++ {
			if i >= n { //剪枝
				break
			}
			list = append(list, i)
			// fmt.Println(list)
			backtracing(n, k, i+1)
			list = list[:len(list)-1]
		}
		return
	}
	backtracing(n, k, 1)
	return res
}

func sum(list []int) int {
	res := 0
	for i := range list {
		res = res + list[i]
	}
	return res
}

func main() {
	n, k := 7, 3
	res := combinationSum3(n, k)
	fmt.Println(res)
}
