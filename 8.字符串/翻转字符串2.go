package main

import "fmt"

func reverse(s string, k int) string {

	return s
}

func reverseStr(s string, k int) string {
	str := []byte(s)
	for i:=0;i<len(str);i=2*k+i {
		if len(str[i:])>=k {
			swap(str[i:k+i])    //切片是开区间  左闭右开
		}else if len(str[i:])<k {
			swap(str[i:])
		}	      
	}
	fmt.Println(string(str))
	return string(str)
}


func swap(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {
	s := "abcdefg"
	val := 2
	reverseStr(s,val)
}
