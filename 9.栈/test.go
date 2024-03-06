package main

import (
	"fmt"
)

func main() {
	ls := []int{1, 2, 3}
	a := 3
	b := 4
	if true && ls[a-b] == 0 {
		fmt.Println("ok")
	}
}
