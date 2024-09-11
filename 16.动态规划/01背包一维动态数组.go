package main

import "fmt"

func bag02(value, weight []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)    //先初始化一下
	for i := 0; i <= bagWeight; i++ {
		if i >= weight[0] {
			dp[i] = value[0]
		}
	}
	fmt.Println(dp)

	for i := 1; i < len(value); i++ {
		for j := 1; j <= bagWeight; j++ {    //注意这里的顺序
			if weight[i] <= j {
				dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			}
		}
		fmt.Println(dp)
	}
	return dp[bagWeight]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	value := []int{2,3 ,1 ,5 ,4 ,3}
	weight := []int{2,2 ,3 ,1 ,5 ,2}
	bag := 1
	fmt.Println(bag02(value, weight, bag))
}
