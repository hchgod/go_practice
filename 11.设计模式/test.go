package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(scanner)
	fmt.Print("请输入您的姓名：")
	scanner.Scan() // 等待用户输入
	fmt.Printf("您输入的姓名是：%s\n", scanner.Text())
}
