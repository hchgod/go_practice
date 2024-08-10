package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	start_x, start_y := 0, 0
	offset := 1
	loop := n / 2
	count := 1
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	for ; loop > 0; loop-- {
		i := start_x
		j := start_y
		for ; j < n-offset; j++ {
			res[i][j] = count
			count++
		}
		for ; i < n-offset; i++ {
			res[i][j] = count
			count++
		}
		for ; j > offset-1; j-- {
			res[i][j] = count
			count++
		}
		for ; i > offset-1; i-- {
			res[i][j] = count
			count++
		}
		start_x++
		start_y++
		offset++
	}
	if n%2 == 1 {
		res[n/2][n/2] = n * n
		return res
	} else {
		return res
	}
}

func main() {
	n := 4
	res := generateMatrix(n)
	fmt.Print(res)
}
