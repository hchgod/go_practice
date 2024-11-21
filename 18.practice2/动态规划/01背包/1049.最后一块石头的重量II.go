package main

import "fmt"

func lastStoneWeightII(stones []int) int {
	half := Sum(stones)/2
	dp := make([]int,half+1)
	for i := 0; i < len(stones); i++{
		for j := half; j > 0; j--{
			if j-stones[i] >= 0{
				dp[j] = max(dp[j],dp[j-stones[i]]+stones[i])
			}
		}
	}
	fmt.Println(dp)
	return Sum(stones) - 2 * dp[half]
}

func Sum(arr []int) int {
	sum := 0
	for i := range arr{
		sum = sum + arr[i]
	}
	return sum
}

func main() {
	lastStoneWeightII([]int{56,89,25})
}