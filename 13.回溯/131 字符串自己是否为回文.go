package main

import (
	"fmt"
)

func partition(candidates string) [][]string {
	if candidates == "" {
		return nil
	}
	result := [][]string{}
	temp := []string{}
	var backtracing func(start int)
	backtracing = func(start int) {
		if start == len(candidates) {
			copyTemp := make([]string, len(temp))
			copy(copyTemp, temp) //深copy
			result = append(result, copyTemp)
			return 
		}
		for i := start; i < len(candidates); i++ {
			str := candidates[start : i+1]              //字符串可以直接切片
			if isPalindrome(str) {
				temp = append(temp, str)
				backtracing(i+1)
				temp = temp[:len(temp)-1]
			}
		}
	}
	backtracing(0)
	return result
}

func isPalindrome(str string) bool {
	for i,j := 0,len(str)-1;i<j;i, j = i+1, j-1{
		if string(str[i]) != string(str[j]){
			return false
		}
	}
	return true
}

func main() {
	a := "aaab"
	res := partition(a)
	fmt.Println(res)
}