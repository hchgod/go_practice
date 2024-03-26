package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	//fmt.Scanf("%d+%d", &num1, &num2)   //格式化输入
	//fmt.Scan(&num1, &num2)   //标准输入
	fmt.Scanln(&num1, &num2)
	fmt.Print(num1, num2)
}
