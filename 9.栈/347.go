package main

import (
	"fmt"
	"sort"
)

// func topKFrequent(nums []int, k int) []int {
// 	hashmap := make(map[int]int)
// 	res := []int{}
// 	for _,value := range nums {
// 		hashmap[value]++
// 	}
// 	for key,_:=range hashmap{
//         res=append(res,key)
// 	}
// 	fmt.Println(hashmap)
// 	fmt.Println(res)
// 	sort.Slice(res,func(a,b int)bool{
//         return hashmap[res[a]]>hashmap[res[b]]
//     })
// 	return res[:k]
// }

func topKFrequent(nums []int, k int) []int {
	hashmap := make(map[int]int)
	for _,value := range nums{
		hashmap[value]++
	}
	res := []int{}
	for key,_ := range hashmap{
		res = append(res, key)
	}
	fmt.Println(res)
	sort.Slice(res,func(a,b int)bool{
		return hashmap[res[a]]>hashmap[res[b]]
	})
	return res[:k]
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 4}
	k := 2
	res := topKFrequent(nums, k)
	fmt.Println(res)
}
