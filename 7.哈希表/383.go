package main

import(
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	res := [26]int{}
	for _ ,v := range magazine{
		res[v-rune('a')]++ 
	}
	for _ ,v := range ransomNote{
		res[v-rune('a')]--
		a := res[v-rune('a')]
		if a<0 {
			return false
		} 
	}
	return true
} 

func main() {
	s := "aa"
	t := "aab"
	res := canConstruct(s,t)
	fmt.Print(res)
}