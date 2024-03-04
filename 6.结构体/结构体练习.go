package main

import(
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}


func main() {
	var stu1 Student
	stu1.Name = "张三"
	stu1.Age = 28
	stu1.Score = 88

	var stu2 Student
	stu2.Name = "李四"
	stu2.Age = 25
	stu2.Score = 100

	stu1.next = &stu2

	var stu3 Student
	stu3.Name = "王五"
	stu3.Age = 18
	stu3.Score = 60

	stu2.next = &stu3

	// print_stu(&stu1)
	var stu *Student = &Student {}
	fmt.Println(stu)
}

func print_stu(temp *Student) {
	for temp != nil {
		fmt.Println(temp)
		fmt.Println(*temp)
		temp = temp.next
	}
}


