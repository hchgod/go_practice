package main

import "fmt"    //嵌套的字典

func main() {
	order1 := map[string]int{
		"宫保鸡丁" : 45,
		"糖醋里脊" : 30,
	}	
	order2 := map[string]int{
		"糖醋里脊": 55,
		"回锅肉":  50,
	}
	order3 := map[string]int{
		"奶茶": 20,
		"啤酒": 10,
	}
	var menu []map[string]int     
	menu = append(menu, order1, order2, order3)
	// fmt.Println(menu)
	for key,value := range menu {
		fmt.Printf("第%d的菜单：\n",key)
		for name,price := range value {
			fmt.Printf("\t菜:%s, 价格:%d\n", name, price)
		}
	}
}