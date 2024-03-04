package main

import "fmt"

func mat(n int,m int) [][]int {
	res := make([][]int,n)
	for i:=0;i<n;i++ {
		res[i] = make([]int, m)
	}
	return res
}

func main() {
	n,m := 3,4
	fmt.Println(mat(n,m))
}