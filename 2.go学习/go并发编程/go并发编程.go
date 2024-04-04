package main

import (
	"fmt"
	//"time"
	"sync"
)
// func print(){
// 	fmt.Println("开始时")
// }

// func main() {
// 	go print()
// 	fmt.Println("你好")
// 	time.Sleep(time.Second*2)
// }


 
var wg sync.WaitGroup
 
func hello(i int) {
    defer wg.Done()  //goroutine结束就登记-1
    fmt.Println("Hello Goroutine!", i)
}
 
func main() {
    for i := 0; i < 10; i++ {
        wg.Add(1)  //计数器，启动一个Goroutine就登记+1，默认值0
        go hello(i)
    }
    wg.Wait()  //阻塞当前协程（这里是main协程），直到计数器数值归零
}