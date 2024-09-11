package main

import (
	"fmt"
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
var instance *ShoppingCartManager //一个类只有一个实例，并提供一个全局访问点来访问这个实例


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
	strs := []string{"hello","niubi","ddkfj"}

	for _,str := range strs {
		if str == "" {
			break
		}
		parts := strings.Fields(str)
		itemName := parts[0]
		quantity := 5
		if len(parts) > 1 {
			fmt.Sscanf(parts[1], "%d", &quantity)
		}

		// 获取购物车实例并添加商品
		cart.addToCart(itemName, quantity)
	}

	// 输出购物车内容
	cart.viewCart()
}