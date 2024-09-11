package main

import "fmt"

// 咖啡接口
type MakeCoffee interface {
	brew()
}

// 具体的黑咖啡类
type BlackCoffee struct{}

func (bc *BlackCoffee) brew() {
	fmt.Println("Brewing Black Coffee")
}

// 具体的拿铁类
type Latte struct{}

func (l *Latte) brew() {
	fmt.Println("Brewing Latte")
}

// 装饰者抽象类
type Coffee struct {
	coffee MakeCoffee
}

func (c *Coffee) brew() {
	c.coffee.brew()
}

// 具体的牛奶装饰者类
type AddMilk struct {
	Coffee
	//毫升
	milliliter float64
}

func (AM *AddMilk) brew() {
	AM.Coffee.brew()
	fmt.Println("Adding Milk",AM.milliliter,"ml")
}

// 具体的糖装饰者类
type AddSugar struct {
	Coffee
}

func (AS *AddSugar) brew() {
	AS.Coffee.brew()
	fmt.Println("Adding Sugar")
}

func main() {
	for i := 0; i < 10; i++ {
		var coffeeType, condimentType int = 2,1

		// 根据输入制作咖啡
		var coffee MakeCoffee
		if coffeeType == 1 {
			coffee = &BlackCoffee{}
		} else if coffeeType == 2 {
			coffee = &Latte{}
		} else {
			fmt.Println("Invalid coffee type")
			continue
		}

		// 根据输入添加调料
		if condimentType == 1 {
			coffee = &AddMilk{Coffee: Coffee{coffee: coffee}}
		} else if condimentType == 2 {
			coffee = &AddSugar{Coffee: Coffee{coffee: coffee}}
		} else {
			fmt.Println("Invalid condiment type")
			continue
		}

		// 输出制作过程
		coffee.brew()
	}
}