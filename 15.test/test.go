package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string) // 创建字符串类型的通道
	ch2 := make(chan string)

	go sendData(ch1) // 启动两个发送数据的 goroutine
	go sendData(ch2)

	for {
		select {
		case msg1 := <-ch1: // 从 ch1 接收数据
			fmt.Println("接收到数据：", msg1)
		case msg2 := <-ch2: // 从 ch2 接收数据
			fmt.Println("接收到数据：", msg2)
		case <-time.After(3 * time.Second): // 定时器，3 秒后超时
			fmt.Println("超时！")
			return
		}
	}
}

// 发送数据到通道
func sendData(ch chan<- string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		ch <- fmt.Sprintf("数据 %d", i) // 发送数据到通道
	}
	close(ch) // 关闭通道
}