package main

import (
	"fmt"
	"strconv"
)

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回
// 这道题是2个不同的集合求组合  
var (
	res []string
	path string
	m []string
)

//字符串类型是不能append，可以使用"" + ""
func letterCombinations(digits string) []string {
	res = []string{}
	m = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	if len(digits) == 0 {
		return res
	}
	backtracing(digits, 0)

	return res
}

func backtracing(digits string, index int) {

	//这是递归边界
	if index == len(digits) {
		res = append(res, path)
		return
	}

	//找到digits的每一个元素
	subindex, _ := strconv.Atoi(string(digits[index]))

	//找到对应的m数组的对应字符串
	for i := 0; i < len(m[subindex]); i++ {
		path = path + string(m[subindex][i])
		backtracing(digits, index+1)
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(letterCombinations("232"))
}
