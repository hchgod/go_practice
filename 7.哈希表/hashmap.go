package main

import(
	"fmt"
)

func isAnagram(s string,t string) bool {
	res := [26]int{}
	for _,v := range s {
		res[v-rune('a')]++
	}
	for _,v := range t {
		res[v-rune('a')]--
	}
	return res == [26]int{}
} 

func main() {
	s := "anagram"
	t := "nagaram"
	res := isAnagram(s,t)
	fmt.Print(res)
}