package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0  //在 Go 语言中，全局变量的定义必须使用 var 关键字，不能使用短变量声明 :=
//锁的作用域：如果将 sync.Mutex 定义为函数内的局部变量，那么每次调用该函数都会创建一个新的互斥锁。这样做会导致每个函数调用都拥有自己的锁实例，无法实现对共享资源的正确保护，因为这些锁实例并不同步。所以，为了确保对共享资源的正确保护，通常需要将互斥锁定义为全局变量或者结构体的字段，以便多个函数能够共享同一个锁实例。
var lock sync.Mutex
//共享资源保护：全局变量或者结构体字段可以跨多个函数或方法，保护涉及共享资源的操作。如果每个函数都有自己的锁实例，那么无法有效地保护全局变量或结构体字段，因为各个函数之间无法共享同一个锁实例。
//锁变量必须是全局变量
func increase(){
	lock.Lock()
	defer lock.Unlock()
	count++
}

func main() {
	for i:=0; i<1000; i++ {
		go increase()
	}
	time.Sleep(time.Second*3)
	fmt.Println("final:",count)
}