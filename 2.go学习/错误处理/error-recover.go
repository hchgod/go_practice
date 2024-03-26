package main

import (
	"fmt"
)
func slicedata(s []string) {
	defer func() {
		err := recover()    //接受到panic报错  并打印panic错误 
		if  err != nil {
			fmt.Printf("有异常：%s", err)
		}
	}()       //这个函数是自执行     并且立即执行
	fmt.Println(s[1])
}


func main() {
	s := []string{"a", "b"}
	slicedata(s)
}