package main

import (
	"fmt"
)

// 本题包含了   树枝和树层去重   一定要注意
var (
	res  [][]int
	path []int
)

func findSubsequences(nums []int) [][]int {
	res = [][]int{}
	path = []int{}
	used := map[int]bool{}
	backtracing(nums, 0, used)
	return res
}

func backtracing(nums []int, startIndex int, used map[int]bool) {

	for i := startIndex; i < len(nums); i++ {
		if (len(path) != 0 && path[len(path)-1] > nums[i]) || used[nums[i]] == true {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		if len(path) > 1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			fmt.Println("res is ", res, "path is", path, "i is", i, "used", used)
		}
		backtracing(nums, i+1, used)
		path = path[:len(path)-1]
		fmt.Println("出来后的值", "res is ", res, "path is", path, "i is", i, "used", used)
	}
}

func main() {
	fmt.Println(findSubsequences([]int{1, 2, 2}))
}
