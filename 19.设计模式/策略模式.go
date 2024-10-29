package main

import "fmt"

// 抽象购物优惠策略接口
type DiscountStrategy interface {
	applyDiscount(originalPrice int) int
}

// 九折优惠策略
type DiscountStrategy1 struct{}

func (d *DiscountStrategy1) applyDiscount(originalPrice int) int {
	return int(float64(originalPrice) * 0.9)
}

// 满减优惠策略
type DiscountStrategy2 struct {
	thresholds []int
	discounts  []int
}

func (d *DiscountStrategy2) applyDiscount(originalPrice int) int {
	for i := len(d.thresholds) - 1; i >= 0; i-- {
		if originalPrice >= d.thresholds[i] {
			return originalPrice - d.discounts[i]
		}
	}
	return originalPrice
}

// 上下文类
type DiscountContext struct {
	discountStrategy DiscountStrategy
}

func (d *DiscountContext) setDiscountStrategy(discountStrategy DiscountStrategy) {
	d.discountStrategy = discountStrategy
}

func (d *DiscountContext) applyDiscount(originalPrice int) int {
	return d.discountStrategy.applyDiscount(originalPrice)
}

func main() {
	// 读取需要计算优惠的次数
	var N int = 3

	for i := 0; i < N; i++ {
		// 读取商品价格和优惠策略
		var M, strategyType int = 10, 2

		// 根据优惠策略设置相应的打折策略
		var discountStrategy DiscountStrategy
		switch strategyType {
		case 1:
			discountStrategy = &DiscountStrategy1{}
		case 2:
			discountStrategy = &DiscountStrategy2{
				thresholds: []int{100, 150, 200, 300},
				discounts:  []int{5, 15, 25, 40},
			}
		default:
			// 处理未知策略类型
			fmt.Println("Unknown strategy type")
			return
		}

		// 设置打折策略
		context := &DiscountContext{}
		context.setDiscountStrategy(discountStrategy)

		// 应用打折策略并输出优惠后的价格
		discountedPrice := context.applyDiscount(M)
		fmt.Println(discountedPrice)
	}
}
