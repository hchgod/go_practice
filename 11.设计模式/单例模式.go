package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// ShoppingCartManager 实现购物车管理
type ShoppingCartManager struct {
	cart   map[string]int
	keys   []string // 用于保存插入顺序的键
	mutext sync.Mutex
}

var once sync.Once
var instance *ShoppingCartManager

// getInstance 获取购物车实例
func getInstance() *ShoppingCartManager {
	once.Do(func() {
		instance = &ShoppingCartManager{
			cart: make(map[string]int),
		}
	})
	return instance
}

// addToCart 将商品添加到购物车
func (scm *ShoppingCartManager) addToCart(itemName string, quantity int) {
	scm.mutext.Lock()
	defer scm.mutext.Unlock()

	// 检查是否已经存在，不存在则添加到 keys 中
	if _, exists := scm.cart[itemName]; !exists {
		scm.keys = append(scm.keys, itemName)
	}

	scm.cart[itemName] += quantity
}

// viewCart 查看购物车内容并按照插入顺序输出
func (scm *ShoppingCartManager) viewCart() {
	scm.mutext.Lock()
	defer scm.mutext.Unlock()

	for _, itemName := range scm.keys {
		quantity := scm.cart[itemName]
		fmt.Printf("%s %d\n", itemName, quantity)
	}
}

func main() {
	cart := getInstance()
	scanner := bufio.NewScanner(os.Stdin)   //创建一个从标准输入读取数据的 Scanner 对象
	//scanner 是一个对象    bufio的作用是减少对IO的压力，输入后先等等  一起IO   减少对底层I/O操作的调用次数
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		parts := strings.Fields(input)
		itemName := parts[0]
		quantity := 0
		if len(parts) > 1 {
			fmt.Sscanf(parts[1], "%d", &quantity)
		}

		// 获取购物车实例并添加商品
		cart.addToCart(itemName, quantity)
	}

	// 输出购物车内容
	cart.viewCart()
}