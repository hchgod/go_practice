package main

import "fmt"

// func main() {
// 	defer fmt.Println("我是第一个被注册的")   // 3
// 	fmt.Println("main函数中调用的Println") // 1
// 	defer fmt.Println("我是第二个被注册的")   // 2
// }

// 无论你在什么地方注册defer语句,它都会在所属函数执行完毕之后才会执行, 并且如果注册了多个defer语句,那么它们会按照后进先出的原则执行



// golang里面有两个保留的函数：
// init函数（能够应用于所有的package）
// main函数（只能应用于package main）
// 这两个函数在定义时不能有任何的参数和返回值


const constValue  = 998 // 1
var gloalVarValue int = abc() // 2
func init() { // 3
 fmt.Println("执行main包中main.go中init函数")
}
func main() { // 4
 fmt.Println("执行main包中main.go中main函数")
}
func abc() int {
 fmt.Println("执行main包中全局变量初始化")
 return 998
}