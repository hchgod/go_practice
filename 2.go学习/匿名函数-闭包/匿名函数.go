package main

import "fmt"

// 方式一
var a = func() {
	fmt.Println("你好")
}

var (
	b = func() {
		fmt.Printf("b")
	}
)

// 匿名函数的写法
// func main() {
// 	func(){
// 		fmt.Printf("你好")
// 	}()
// }

// func main() {
// 	a := func(s string) {
// 		fmt.Println(s)
// 	}
// 	a("hello lnj")
// }

func main() {
	test(func(s string) {
		fmt.Println(s)
	})
}

func test(f func(s string)) {
	f("hello lnj")
}
