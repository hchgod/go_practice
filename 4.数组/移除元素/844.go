package main

import (
	"fmt"
)

func backspaceCompare(s, t string) bool {
	s_index, t_index := len(s)-1, len(t)-1
	s_skip, t_skip := 0, 0
	for s_index >= 0 || t_index >= 0 {
		for s_index >= 0 {
			if s[s_index] == '#' {
				s_skip++
				s_index--
			} else if s_skip > 0 {
				s_skip--
				s_index--
			} else {
				break
			}
		}
		for t_index >= 0 {
			if s[t_index] == '#' {
				t_skip++
				t_index--
			} else if t_skip > 0 {
				t_skip--
				t_index--
			} else {
				break
			}
		}
		if s_index >= 0 && t_index >= 0 {
			if s[s_index] != t[t_index]{
				return false
			} 
		} else if s_index >= 0 || t_index >= 0 {
			return false
		}
		s_index--
		t_index--
	}
	return true
}

func main() {
	s := "ad#c"
	t := "ab#c"
	res := backspaceCompare(s, t)
	fmt.Print(res)
}
