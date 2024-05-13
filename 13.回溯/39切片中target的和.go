package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	fmt.Println(candidates)
	if candidates == nil {
		return nil
	}
	result := [][]int{}
	temp := []int{}
	var backtracing func(start int)
	backtracing = func(start int) {
		if sum(temp) == target {
			copyTemp := make([]int, len(temp))
			copy(copyTemp, temp) //深copy
			result = append(result, copyTemp)
			return
		}
		if sum(temp) > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			temp = append(temp, candidates[i])
			backtracing(i)
			temp = temp[:len(temp)-1]
		}
	}
	backtracing(0)

	return result
}

func sum(ls []int) int {
	sum := 0
	for i := 0; i < len(ls); i++ {
		sum += ls[i]
	}
	return sum
}

func main() {
	a := []int{2, 3, 6, 7}
	res := combinationSum(a, 7)
	fmt.Println(res)
}
