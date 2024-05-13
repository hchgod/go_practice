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
		if isPalindrome(fmt.Sprint(temp)) == true {
			copyTemp := make([]string, len(temp))
			copy(copyTemp, temp) //深copy
			fmt.Println("copyTemp:",copyTemp)
			result = append(result, copyTemp)
			return
		}
		for i := start; i < len(candidates); i++ {
			temp = append(temp, string(candidates[i]))
			backtracing(i+1)
			temp = temp[:len(temp)-1]
		}
	}
	backtracing(0)
	return result
}

func isPalindrome(str string) bool {
	fmt.Println("str is ",str)
	for i,j := 0,len(str)-1;i<j;i, j = i+1, j-1{
		if string(str[i]) != string(str[j]){
			fmt.Println(false)
			return false
		}
	}
	fmt.Println(true)
	return true
}

func main() {
	a := "aab"
	res := partition(a)
	fmt.Println(res)
}