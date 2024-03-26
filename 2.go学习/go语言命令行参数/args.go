package main

import (
	"os"
	"fmt"
)

func main() {
	nums := len(os.Args)
	for i := 0; i < nums; i++ {
		fmt.Println(os.Args[i])
	}
}
