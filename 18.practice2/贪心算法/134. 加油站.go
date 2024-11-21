package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	cursum := gas[0] - cost[0]
	start := 0
	for j := 1; j < 2*len(gas); j++ {
		i := j % len(gas)
		if start == i && cursum >= 0{
			fmt.Println("xxx", "start；", start, "cursum：", cursum)
			fmt.Println(start)
			return start
		}
		if cursum < 0 {
			start = i
			cursum = gas[i] - cost[i]
			fmt.Println("ttttt","start；", start, "cursum：", cursum, "i is", i)
			continue
		}
		cursum = cursum + gas[i] - cost[i]
		fmt.Println("wwww","start；", start, "cursum：", cursum, "i is", i)
	}
	fmt.Println(-1)
	return -1
}

func main() {
	canCompleteCircuit([]int{4,5,3,1,4}, []int{5,4,3,4,2})
	canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})
	canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
}
