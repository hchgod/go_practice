package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
    hashmap := make(map[int]int)
	for index,val := range nums{
		hashmap[val] = index+1
	}
	for index,val := range nums{
		diff := target - val
		if hashmap[diff]!=0 {
			if index == hashmap[diff]-1{
				continue
			}else{
				return []int{index, hashmap[diff]-1}
			}
		}else{
			continue
		}
	}
	return []int{}
}


func main() {
	n := []int{2, 18, 11, 15,1}
	target := 3
	res := twoSum(n,target)
	fmt.Print(res)
}
