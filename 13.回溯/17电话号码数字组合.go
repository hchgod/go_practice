package main

import (
	"fmt"
	"strings"
)

func letterCombinations(digits string) []string {
	var result []string
	m := []string{"", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var tmp []string
	var backtracing func( start int)
	backtracing = func(start int) {
		fmt.Println(len(tmp))
		if len(tmp) == len(digits) {
			result = append(result, strings.Join(tmp,""))
			return   //注意这里要return，否则会一直递归下去
		}
		letter := m[int(digits[start]-'0')]
		fmt.Println(letter)
		for i := 0; i < len(letter); i++ {
			fmt.Println(letter[i])
			tmp = append(tmp, string(letter[i]))
			backtracing(start+1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	backtracing(0)
	return result
}

func main() {
	res := letterCombinations("23")
	fmt.Println(res)
}
