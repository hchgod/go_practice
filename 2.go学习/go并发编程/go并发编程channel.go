package main

import (
	"fmt"
	"reflect"
	"time"
)

// 程序1
func makeBuns(filling string, buns chan string) {
	fmt.Printf("开始做%s馅的包子\n", filling)
	// 1. 剁馅
	fmt.Printf("开始剁%s馅...\n", filling)
	// 2. 擀皮
	fmt.Println("开始擀皮...")
	time.Sleep(time.Second)
	// 3. 包包子
	fmt.Printf("开始包%s馅的包子...\n", filling)
	// 4. 蒸
	fmt.Printf("开始蒸%s馅的包子...\n", filling)
	// 5. 蒸好了
	time.Sleep(time.Second * 1)
	fmt.Printf("%s馅包子已经蒸好了，可以上菜了,蒸好的时间是:%s\n", filling, time.Now())
	// 6. 上菜
	// 在这个位置把蒸好的包子放到通道内
	buns <- filling
	fmt.Printf("%s馅包子已经放在了上菜区，放置时间是:%s\n", filling, time.Now())
}

func main() {
	buns := make(chan string, 5) //创建一个channl
	defer close(buns)
	filling := []string{"韭菜", "鸡蛋", "西葫芦"}
	for _, value := range filling {
		go makeBuns(value, buns)
	}

	for i := 0; i < len(filling); i++ {
		bun := <-buns //读数据   从管道中读数据
		fmt.Println(reflect.TypeOf(bun))
		fmt.Printf("上菜：%s,上菜时间：%s \n", bun, time.Now())
		time.Sleep(time.Second * 3)
	}
}
