package main

import (
	"fmt"
	"reflect"
)

type EmptyInterface interface{}

func main() {
	var i EmptyInterface
	i = "这是一个字符串"
	fmt.Println("此时的空接口的值是:", i)

	v2 := 7788
	i = v2
	fmt.Println("此时的空接口的值是:", i)
	// i.(int)  是类型断言
	s1, ok := i.(string) //   返回的是这个空接口的值和是否对应的类型
	fmt.Println(reflect.TypeOf(s1))
	fmt.Println(s1)
	fmt.Println(ok)
}
