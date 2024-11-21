package main

import (
	"fmt"
	"sort"
)

var (
	res [][]int
	tmp []int
)

// 本题总结：
// 1.  确定递归边界  什么时候退出
// 2. 确定递归传的参数， 并且 注意参数 index的起始位置  。
// 3. 树层去重   树枝也重复
func combinationSum2(candidates []int, target int) [][]int {
	res = [][]int{}
	tmp = []int{}
	sum := 0
	sort.Ints(candidates)
	backtracing(candidates, target, sum, 0)

	return res
}

func backtracing(candidates []int, target int, sum int, startIndex int) {
	if sum > target {
		return
	}
	if sum == target {
		res1 := make([]int, len(tmp))
		copy(res1, tmp)
		res = append(res, res1)
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		if (i == 0 && candidates[i] <= target - sum) || (candidates[i] <= target - sum && candidates[i] != candidates[i-1]){
			tmp = append(tmp, candidates[i])
			sum = sum + candidates[i]
			backtracing(candidates, target, sum, i+1)
			sum = sum - tmp[len(tmp)-1]
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func main() {
	fmt.Println(combinationSum2([]int{2,5,2,1,2}, 5))
}




