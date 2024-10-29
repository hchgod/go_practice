package main

import "fmt"

func function(n int) {
	stone := []int{31,26,33,21,40}
	half := Sum(stone)/2
	dp := make([]int,half+1)
	for i := 0; i < 5; i++{
		for j := half; j > 0; j--{
			if j-stone[i] >= 0{
				dp[j] = max(dp[j],dp[j-stone[i]]+stone[i])
			}
		}
		fmt.Println(dp[35:])
	}
	fmt.Println(dp)
	return
}

func main() {
	function(5)
}

func Sum(arr []int) int {
	sum := 0
	for i := range arr{
		sum = sum + arr[i]
	}
	return sum
}