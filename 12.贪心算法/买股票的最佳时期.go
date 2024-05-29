package main

import "fmt"

func maxProfit(prices []int) int {
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		diff := prices[i+1] - prices[i]
		if diff > 0 {
			res = res + diff
		}
	}
	return res
}

func main() {
	a := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(a)
	fmt.Println(res)
}
