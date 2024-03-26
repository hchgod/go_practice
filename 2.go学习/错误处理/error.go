package main

import (
	"errors"
	"fmt"
)

func division(a, b float64) (res float64,err error) {
	fmt.Println("计算的值为%d,%d",a,b)
	if a == 0{
		return 0, errors.New("这是自定义错误  分母不能为0")
	}else{
		return b/a,nil
	}
}

func main() {
	err := errors.New("这是自定义错误")
	fmt.Println(err)
	err2 := fmt.Errorf("这是自定义的错误对象%s",":  这是错误内容")
	//fmt.Errorf() 函数用于创建一个新的错误值，并返回一个实现了error接口的错误对象。
	//格式化并返回新的error  
	fmt.Println("这是fmt定义的错误对象",err2.Error())
}
