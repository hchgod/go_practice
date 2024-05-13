package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	candidates = removeDuplicates(candidates)
	sort.Ints(candidates)
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
			backtracing(i + 1)
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

func removeDuplicates(list []int) []int {
	res := []int{}
	map1 := make(map[int]bool)
	for i := 0; i < len(list); i++ {
		value := map1[list[i]]
		if !value {
			map1[list[i]] = true
			res = append(res, list[i])
		}else {
			continue
		}
	}
	return res
}

func main() {
	a := []int{10,1,2,7,6,1,5}
	res := combinationSum(a, 8)
	fmt.Println(res)
}