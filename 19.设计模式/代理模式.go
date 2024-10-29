package main

import "fmt"

// 抽象主题
type HomePurchase interface {
	requestPurchase(area int)
}

// 真实主题
type HomeBuyer struct{}

func (hb *HomeBuyer) requestPurchase(area int) {
	fmt.Println("YES")
}

// 代理类
type HomeAgentProxy struct {
	homeBuyer HomePurchase
}

func (hap *HomeAgentProxy) requestPurchase(area int) {
	if area > 100 {
		hap.homeBuyer.requestPurchase(area)
	} else {
		fmt.Println("NO")
	}
}

func main() {
	var buyerProxy HomePurchase = &HomeAgentProxy{homeBuyer: &HomeBuyer{}}

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var area int
		fmt.Scan(&area)
		buyerProxy.requestPurchase(area)
	}
}