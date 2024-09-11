package main

import "fmt"

// Animal 是一个接口，定义了动物的行为
type Animal interface {
    Speak() string
}

// Dog 是 Animal 的具体实现
type Dog struct{}

func (d *Dog) Speak() string {
    return "开始叫了,我是狗"
}

// Cat 是 Animal 的具体实现
type Cat struct{}

func (c *Cat) Speak() string {
    return "开始叫了,我是猫"
}

// AnimalFactory 是一个工厂函数，创建 Animal 对象
func AnimalFactory(animalType string) Animal {
    if animalType == "dog" {
        return &Dog{}
    } else if animalType == "cat" {
        return &Cat{}
    }
    return nil
}

func main() {
    // 使用工厂函数创建 Dog 对象
	var animal Animal
    animal = AnimalFactory("dog")
    fmt.Println(animal.Speak())

    // 使用工厂函数创建 Cat 对象
    animal = AnimalFactory("cat")
    fmt.Println(animal.Speak())
}
