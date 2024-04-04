package main

import(
	"fmt"
	"time"
)

func main() {
	//第一种
	var a,b int = 1,2
	fmt.Println(a,b)
	var i,j = 1,2
	fmt.Println(i,j)
	//第二种
	var c,d int
	c,d = 3,4
	fmt.Println(c,d)
	var (
		e = 5
		f = 6
	)
	fmt.Println(e,f)
	var g, h string = "123", "hello"
	fmt.Println(g,h)
	time.Sleep(time.Second*10)
}


