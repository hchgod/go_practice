package main
import (
	"fmt"
)

func gettype(i interface{}) {
	
	//存在一个小错误。在 Go 语言中.(type) 是用在 switch 语句中的特殊写法，用于获取接口值的动态类型。但是，在使用 .(type) 时，它必须在一个 switch 语句的 case 中使用，而不是在单独的变量赋值语句中。
	switch t := i.(type) {
	case int :
		fmt.Println("这是int类型",t)
	case string :
		fmt.Println("这是string类型")
	case float64 :
		fmt.Println("这是float64类型")
	default:
		fmt.Println("未知类型")
	}
}

func main() {
	var nili interface{}
	gettype(nili)

}