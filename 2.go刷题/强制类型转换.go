package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main() {
	s1 := "50"
	num1, _:= strconv.Atoi(s1)
	fmt.Print(reflect.TypeOf(num1))
	i := 7275
	sn := strconv.Itoa(i)
	fmt.Println(sn)
	fmt.Println("类型:", reflect.TypeOf(sn))
}