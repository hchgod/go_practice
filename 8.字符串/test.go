package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	s := strings.Split("a,b,c", ",")
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i])
	}
}
