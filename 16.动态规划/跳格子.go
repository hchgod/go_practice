package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	res := 0
	length := len(cost)
	if length == 0 {
		return res
	}

	var index int = 0
	for index+2 < length {
		mid := index
		front := index + 1
		if front == length-1 {
			return res + cost[front]
		}
		if cost[mid] >= cost[front] {
			res = res + cost[front]
			index = front
		} else {
			res = res + cost[mid]
			index = mid
		}
	}
	return res
}

func main() {
	//var a int = 3
	arr := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(arr))
}
