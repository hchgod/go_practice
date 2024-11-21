package main

import (
	"fmt"
)
// 本题总结：
// 1. 遍历的是一个数组   1~n
// 2. 需要start index
// 3. 确定递归深度  为k    就是递归的边界  确定退出递归的时机   比如sum   比如最后一个值
// 4. 确定递归传的参数
// 5. 这道题结果集 res 中是叶子节点  所以要判断  加入叶子节点的时机
var (
	res [][]int
	tmp []int
)
func combine(n int, k int) [][]int {
	tmp, res = []int{}, make([][]int, 0)
	dfs(1, n, k)
	return res
}

func dfs(startindex int, n int, k int) {
	//加入叶子节点  可以是 sum  
	if len(tmp) == k {
		res_tmp := make([]int, len(tmp))
		copy(res_tmp, tmp)
		res = append(res, res_tmp)
		fmt.Println(res)
		return
	}
	for i := startindex; i <= n; i++ {
		if n-i+1 < k-len(tmp){
			return
		}
		if len(tmp) < k {
			tmp = append(tmp, i)
			dfs(i+1, n, k)
			tmp = tmp[:len(tmp)-1]
		}
	}
	return
}

func main() {
	combine(4,2)
}
