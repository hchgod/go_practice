package main

import "fmt"

// 定义一个泛型结构体
type Pair[T any] struct {
    First  T
    Second T
}

func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

func main() {
    // 使用泛型结构体
    intPair := Pair[int]{First: 1, Second: 2}
    stringPair := Pair[string]{First: "hello", Second: "world"}

    fmt.Println(intPair)
    fmt.Println(stringPair)

	intSlice := []int{1, 2, 3, 4, 5}
    stringSlice := []string{"hello", "world"}

    PrintSlice(intSlice)
    PrintSlice(stringSlice)
}
