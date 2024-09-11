package main

import (
	"fmt"
)

// 抽象原型类
type Prototype interface {
	clone() Prototype
	getDetails() string
}

// 具体矩形原型类
type RectanglePrototype struct {
	color  string
	width  int
	height int
}

// 构造方法
func NewRectanglePrototype(color string, width, height int) *RectanglePrototype {
	return &RectanglePrototype{
		color:  color,
		width:  width,
		height: height,
	}
}

// 实现 Prototype 接口的 clone 方法
func (r *RectanglePrototype) clone() Prototype {
	return &RectanglePrototype{
		color:  r.color,
		width:  r.width,
		height: r.height,
	}
}

// 获取矩形的详细信息
func (r *RectanglePrototype) getDetails() string {
	return fmt.Sprintf("Color: %s, Width: %d, Height: %d", r.color, r.width, r.height)
}

// 客户端程序
func main() {
	// 读取需要创建的矩形数量
	var N int = 3

	// 读取每个矩形的属性信息并创建矩形对象
	for i := 0; i < N; i++ {
		var color string = "Red"
		var width, height int = 3 , 4

		// 创建原型对象
		originalRectangle := NewRectanglePrototype(color, width, height)

		// 克隆对象并输出详细信息
		clonedRectangle := originalRectangle.clone()
		fmt.Println(clonedRectangle.getDetails())
	}
}