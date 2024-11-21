package main

import (
	"fmt"
	"strconv"
)


//这道题太难了  没做出来

var (
	res  []string
	path []string
)

func restoreIpAddresses(s string) []string {
	res = []string{}
	path = []string{}
	backtracing(s, 0)
	return res
}

func backtracing(s string, startindex int) {
	if startindex == len(s) && len(path) == 4 {
		res = append(res, combina(path))
		return
	}

	for i := startindex; i < len(s); i++ {
		str := s[startindex : i+1]
		fmt.Println("str", str)
		if len(path) > 3 || len(str) > 3{
			return
		}
		fmt.Println("str:", str, "结果是：", isIp(str))
		if isIp(str) {
			path = append(path, str)
			fmt.Println("添加str之前的path", path, "str", str, "i is", i, "start is", startindex)
			backtracing(s, i+1)
			path = path[:len(path)-1]
			fmt.Println("截断后的path", path, "str", str, "i is", i, "start is", startindex)
		}
	}
}

func combina(strs []string) string {
	sum := ""
	for _, str := range strs {
		sum = sum + str + "."
	}
	return sum[:len(sum)-1]

}

func isIp(arr string) bool {
	num_0, _ := strconv.Atoi(string(arr[0]))
	if len(path) == 0 && arr == "0" {
		return false
	}
	if (len(arr) != 1 && num_0 <= 0) || len(arr) > 3{
		return false
	}
	num, _ := strconv.Atoi(arr)
	if num > 255 || num < 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(restoreIpAddresses("0000"))
}
