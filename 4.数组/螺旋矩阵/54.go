package main

import "fmt"

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func generateMatrix2(n int, m int) [][]int {
	res := make([][]int, n)
	start_x, start_y := 0, 0
	offset := 1
	loop := min(m, n) / 2
	count := 1
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}
	for ; loop > 0; loop-- {
		i := start_x
		j := start_y
		for ; j < m-offset; j++ {
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
	if min(m, n)%2 == 1 {
		loop := min(m, n) / 2
		for i := 0; (i+loop) < (m - loop); i++ {
			res[n/2][loop+i] = count
			count++
		}
		return res
	} else {
		return res
	}
}

func main() {
	n, m := 3, 4
	res := generateMatrix2(n, m)
	fmt.Print(res)
}
