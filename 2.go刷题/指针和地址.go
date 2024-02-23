package main

import (
	"fmt"
)

type People struct {
	name string
	age int
	sex string
	address string
}

func (p *People) get_info() string {
	return fmt.Sprintf("您的信息如下: \n \t姓名: %s\n \t性别: %s\n \t年龄: %d\n \t地址: %s\n", p.name, p.sex, p.age, p.address)
}

func main() {
	p := People{"Dotbalo", 19,"Man","Russia"}
	p.get_info()
}



