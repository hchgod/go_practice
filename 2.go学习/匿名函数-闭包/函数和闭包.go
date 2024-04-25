package main
import(
	"fmt"
)
// 程序1  返回值
func sum() int { // 只有一个返回值时,返回值列表的()可以省略
	return 1 + 1
}

func main() {
	sum()
}

// 值传递和引用传递
// 程序2 
// 值类型包括基本数据类型（如int、float、bool、string）和结构体（struct）
// 引用类型包括切片（slice）、映射（map）、通道（channel）、接口（interface）、指针（pointer）
// 看这篇博客  https://mp.weixin.qq.com/s/LqEglWaYksPtkszCCPl8xA
