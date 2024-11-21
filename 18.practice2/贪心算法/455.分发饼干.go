package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}
    sort.Ints(g)
	sort.Ints(s)
	index := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if index < len(g) && s[i] >= g[index] {
			res++
			index++
			continue
		}
	}
	fmt.Println(res)
	return res
}

func main(){
	findContentChildren([]int{1,2}, []int{1,2,3})
}