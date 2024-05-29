package main

import "fmt"

func maxSubArray(nums []int) int {
    max := nums[0]
    temp := 0

    for i := 0; i < len(nums); i++{
        temp += nums[i]
        if temp > max{
            max = temp
        }
        if temp < 0 {
            temp = 0
        }
    } 
    return max
}

func main() {
	maxSubArray([]int{-1, 1, 2, 1})
}
