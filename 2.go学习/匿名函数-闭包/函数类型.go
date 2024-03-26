package main

import "fmt"

// func main() {
// 	var sum func(a, b int) int //必须和下面的函数定义是一样的
// 	// 函数类型是引用类型
// 	sum = func(a, b int) int {
// 		return a + b
// 	}
// 	res := sum(3, 4)
// 	fmt.Print(res)
// }

func main() {
	a := func1()
	res := a(3, 4)
	fmt.Println(res)
}

func func1() func(int, int) int { //返回值类型定必须一样    都是func类型
	return func(a, b int) int {
		return a * b
	}
}
