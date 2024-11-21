package main

import (
	"fmt"
	"sort"
)

// 本题包含了   树枝和树层去重   一定要注意
var (
	res  [][]int
	path []int
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res = [][]int{{}}
	path = []int{}
	backtracing(nums, 0)
	return res
}

func backtracing(nums []int, startIndex int) {
	used := make([]bool, len(nums))
	for i := startIndex; i < len(nums); i++ {
		// 注意这里的每个树节点  都需要append到res数组中
		if i != 0 && used[i-1] == true && nums[i-1] == nums[i]{
			used[i] = true
			continue
		}
		path = append(path,nums[i])
		used[i] = true
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		fmt.Println("res is ", res, "path is", path, "i is", i, "used", used)
		backtracing(nums, i+1)
		path = path[:len(path)-1]
		fmt.Println("出来后的值", "res is ", res, "path is", path, "i is", i, "used", used)
	}
}

func main() {
	fmt.Println(subsetsWithDup([]int{5,5,5,5,5}))
}
