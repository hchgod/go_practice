package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	child := 0
	for i, j := 0, 0; i < len(s) && j < len(g); i++ {
		if s[i] < g[j] {
			continue
		} else if s[i] >= g[j] {
			child++
			j++
		}
	}
	return child
}

func main() {
	g := []int{1, 2}
	s := []int{1, 1, 3}
	res := findContentChildren(g, s)
	fmt.Println(res)
}
