package main

import (
	"fmt"
	"sort"
)
type person struct {
	name string
	age int
}

func main() {
	person := []person{
		{"a", 18},
		{"b", 20},
		{"c", 16},
	}
	sort.Slice(person, func(i,j int)bool{
		return person[i].age > person[j].age
	})
	fmt.Println(person)
}
