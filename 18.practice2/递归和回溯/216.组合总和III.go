package main

import (
	"fmt"
)

var (
	res [][]int
	tmp []int
)

func combinationSum3(k int, n int) [][]int {
	tmp, res = make([]int, 0, k), make([][]int, 0)
	dfs(1, k, n, 0)
	return res
}

// 和为n的k个数的组合
func dfs(startindex int, k int, n int, sum int) {
	if sum == n && len(tmp) == k {
		subres := make([]int, k)
		copy(subres, tmp)
		res = append(res, subres)
		return
	}
	for i := startindex; i <= 9; i++ {
		if len(tmp) < k && i <= n && n - sum >= i{
			tmp = append(tmp, i)
			sum = sum + i
			//  因为不能重复  所以  下一层递归的起始位置在 i + 1
			dfs(i+1, k, n, sum)
			sum = sum - tmp[len(tmp)-1]
			tmp = tmp[:len(tmp)-1]
		}
	}
	return
}

func main() {
	fmt.Println(combinationSum3(3, 9))
}
