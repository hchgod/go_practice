package main

import (
	"fmt"
)
//  第一种   匿名函数也是一种闭包     函数内部定义函数
// func main() {
// 	sum := sum()
// 	res := sum(3,4)
// 	fmt.Print(res)
// }

// func sum() func(int,int) int{
// 	return func(a,b int) int{
// 		return a + b
// 	}
// }

// 第二种  main函数内部定义闭包
func main() {
	func (a,b int){
		fmt.Println(a + b) 
	}(3,4)
}

