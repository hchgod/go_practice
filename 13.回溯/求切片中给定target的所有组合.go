package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	fmt.Println(candidates)
	res := [][]int{}
	tmp := []int{}
	var backtracing func(start int)
	backtracing = func(start int) {
		if sum(tmp) == target {
			copytmp := make([]int, len(tmp))
			copy(copytmp, tmp)
			res = append(res, copytmp)
			return
		}
		if sum(tmp) > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			tmp = append(tmp, candidates[i])
			backtracing(i+1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	backtracing(0)
	res = removeDuplicates(res)
	return res
}

func sum(ls []int) int {
	sum := 0
	for i := 0; i < len(ls); i++ {
		sum += ls[i]
	}
	return sum
}
func removeDuplicates(res [][]int) [][]int {
	result := [][]int{}
	m := make(map[string]bool, 0)
	for _,value := range res {
		tmp := fmt.Sprint(value) 
		if !m[tmp] {
			m[tmp] = true
			result = append(result, value)
		}else {
			continue
		}
	}
	return result
}


func main() {
	a := []int{10, 1, 2, 7, 6, 1, 5}
	res := combinationSum(a, 8)
	fmt.Println(res)
}
