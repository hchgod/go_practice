package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
    path []string
    res  []string
)
var layer int=0


func restoreIpAddresses(s string) []string {
    path, res = make([]string, 0), make([]string, 0)
    dfs(s, 0)
    return res
}
func dfs(s string, start int) {  
    fmt.Println("输出res是:",res)
    fmt.Println("输出path是:",path)
    layer++
    fmt.Println("这是第%d层递归",layer)
    if len(path) == 4 {    // 够四段后就不再继续往下递归
        if start == len(s) {      
            str := strings.Join(path, ".")
            res = append(res, str)
        }
        return 
    }
    for i := start; i < len(s); i++ {
        if i != start && s[start] == '0' { // 含有前导 0，无效
            break
        }
        str := s[start : i+1]
        num, _ := strconv.Atoi(str)
        if num >= 0 && num <= 255 {
            path = append(path, str)  // 符合条件的就进入下一层
            dfs(s, i+1)
            path = path[:len(path) - 1]
        } else {   // 如果不满足条件，再往后也不可能满足条件，直接退出
            break
        }
    }
}

func main() {
	a := "235161212131"
	res := restoreIpAddresses(a)
	fmt.Println(res)
}