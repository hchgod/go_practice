package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	if s == "" {
		return nil
	}
	result := []string{}
	temp := []string{}
	var backtracing func(start int)
	backtracing = func(start int) {
		if start == len(s) && len(temp) == 4 {
			copyTemp := strings.Join(temp, ".")
			result = append(result, copyTemp)
			return
		}
		for i := start; i < len(s); i++ {
			str := s[start : i+1] //字符串可以直接切片
			if start != len(s) && len(temp) == 4 {
				return
			}
			if isValid_ip(str, temp) {
				temp = append(temp, str)
				backtracing(i + 1)
				temp = temp[:len(temp)-1]
			}
		}
	}
	backtracing(0)
	return result
}

func isValid_ip(str string, temp []string) bool {
	if str[0] == byte('0') && len(str) != 1 {
		//以0开头的   数字 pass掉
		return false
	}
	value, _ := strconv.Atoi(str)
	// fmt.Println("value is :",value)
	if value > 255 || value < 0 {
		return false
	}
	return true
}

func main() {
	a := "101023"
	res := restoreIpAddresses(a)
	fmt.Println(res)
}
