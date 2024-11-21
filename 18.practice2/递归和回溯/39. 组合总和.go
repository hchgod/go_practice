package main

import (
	"fmt"
)

var (
	res [][]int
	tmp []int
)

// 本题总结：
// 1.  确定递归边界  什么时候退出
// 2. 确定递归传的参数， 并且 注意参数 index的起始位置, 因为元素是否可以重复取   决定了index的起始位置。
// 3. 本题可以无限制的取   所以 递归传参是 i
func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}
	tmp = []int{}
	sum := 0
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
		if candidates[i] <= target - sum{
			tmp = append(tmp, candidates[i])
			sum = sum + candidates[i]
			//  这个地方是 i
			backtracing(candidates, target, sum, i)
			sum = sum - tmp[len(tmp)-1]
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func main() {
	fmt.Println(combinationSum([]int{2,3,6,7}, 7))
}




