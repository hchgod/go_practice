package main

import "fmt"

// func func1(n int) int {
// 	res := 0
// 	return res
// }

func main() {
	switch num := 3; num {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("不是1,2,3,4")
	}
	fmt.Println("程序结束")
}
