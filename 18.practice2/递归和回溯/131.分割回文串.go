package main

import (
	"fmt"
)

var (
	res [][]string
	path []string
)

// 本题总结：
// 1.  每一层要遍历的时候  注意横向的添加


func partition(s string) [][]string {
	res = [][]string{}
	path = []string{}
	backtracing(s, 0)

	return res
}

func backtracing(s string, startIndex int) {
	if startIndex == len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := startIndex; i < len(s); i++ {
		//注意这一行的 分割 相当于添加   并且是回文了话   上一层直接阶段
		str := s[startIndex : i+1]
		// 不是回文就继续 i++   是回文 就进入 if   然后下一层递归的时候把 i + 1
		if isPalindrome(str) {
			fmt.Println("str是回文:", str)
			path = append(path, str)
			fmt.Println("前的path:", path, "i is", i, "startindex is", startIndex)
			//  这个地方注意   如果是回文 就把 i + 1
			backtracing(s, i+1)
			path = path[:len(path)-1]
			fmt.Println("弹出后的path:", path)
		}
	}
}

func isPalindrome(arr string) bool {
	i := 0
	for j := len(arr) - 1; i <= j; j--{
		if arr[i] != arr[j] {
			return false
		}
		i++
	}
	return true

}

func main() {
	fmt.Println(partition("aab"))
}




