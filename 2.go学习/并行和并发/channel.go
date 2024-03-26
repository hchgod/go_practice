package main

func main() {
	var mych chan int
	// Channel声明和初始化
	// 声明: var 变量名chan 数据类型
	// 初始化: mych := make(chan 数据类型, 容量)
	// Channel和切片还有字典一样, 必须make之后才能使用
	// Channel和切片还有字典一样, 是引用类型
	mych = make(chan int, 5)
	list := []string{"mych", "mych","my"}
	for value := range list {
		
	}
}
