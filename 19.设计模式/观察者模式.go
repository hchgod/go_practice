package main

import (
	"fmt"
)

// 观察者接口
type Observer interface {
	Update(hour int)
}

// 主题接口
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// 具体主题实现
type Clock struct {
	observers []Observer
	hour      int
}

func (c *Clock) RegisterObserver(observer Observer) {
	c.observers = append(c.observers, observer)
}

func (c *Clock) RemoveObserver(observer Observer) {
	for i, obs := range c.observers {
		if obs == observer {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			break
		}
	}
}

func (c *Clock) NotifyObservers() {
	for _, observer := range c.observers {
		observer.Update(c.hour)
	}
}


func (c *Clock) Tick() {
	c.hour = (c.hour + 1) % 24 // 模拟时间的推移
	c.NotifyObservers()
}

// 具体观察者实现
type Student struct {
	name string
}

func NewStudent(name string) *Student {
	return &Student{name: name}
}

func (s *Student) Update(hour int) {
	fmt.Println(s.name, hour)
}



func main() {
	// 读取学生数量
	var N int = 2

	// 创建时钟
	clock := &Clock{}

	// 注册学生观察者
	for i := 0; i < N; i++ {
		var studentName string = "hch"
		clock.RegisterObserver(NewStudent(studentName))
	}

	// 读取时钟更新次数
	var updates int = 3

	// 模拟时钟每隔一个小时更新一次
	for i := 0; i < updates; i++ {
		clock.Tick()
	}
}