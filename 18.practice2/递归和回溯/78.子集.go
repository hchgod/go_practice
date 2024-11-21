package main

import (
	"fmt"
)

var (
	res  [][]int
	path []int
)

func subsets(nums []int) [][]int {
	res = [][]int{{}}
	path = []int{}
	backtracing(nums, 0)
	return res
}

func backtracing(nums []int, startIndex int) {
	for i := startIndex; i < len(nums); i++ {
		// 注意这里的每个树节点  都需要append到res数组中
		path = append(path,nums[i])
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		fmt.Println("res is ", res, "path is", path, "i is", i)
		backtracing(nums, i+1)
		path = path[:len(path)-1]
		fmt.Println("出来后的值", "res is ", res, "path is", path, "i is", i)
	}
}

func main() {
	fmt.Println(subsets([]int{0}))
}
