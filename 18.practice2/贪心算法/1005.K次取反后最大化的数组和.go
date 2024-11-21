package main

import (
	"fmt"
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	res := 0
	long := len(nums)
	sort.Ints(nums)
	for i := 0; i < long && k > 0; i++ {
		fmt.Println("xxxqqq",nums,k,nums[i])
		if nums[i] < 0{

			nums[i] = -nums[i]
			k--
			fmt.Println("xxx",nums,k,nums[i])
			continue
		}else if nums[i] == 0 {
			k = 0
			break
		}
	}
	sort.Ints(nums)
	nums[0] = int(math.Pow(-1, float64(k%2))) * nums[0]

	fmt.Println("xxx",nums)
	res = sum(nums)
	return res
}

func sum(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res = res + nums[i]
	}
	fmt.Println(res)
	return res
}

func main(){
	largestSumAfterKNegations([]int{2,-3,-1,5,-4},2)
}