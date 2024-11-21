package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n < 1 {
		return n
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	fmt.Println(points)
	res := 1
	board := points[0][1]
	for i := 1; i < n; i++ {
		if points[i][0] > points[i-1][1] || points[i][0] > board {
			res++
			board = points[i][1]
		}
		fmt.Println("i is ",points[i],"res is ",res)
	}
	return res
}

func main() {
	fmt.Println(findMinArrowShots([][]int{{3, 9},
        {7, 12},
        {3, 8},
        {6, 8},
        {9, 10},
        {2, 9},
        {0, 9},
        {3, 9},
        {0, 6},
        {2, 8},}))
}