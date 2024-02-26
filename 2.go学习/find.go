package main

import (
	"fmt"
)

func binary_find(arr []int, num int) (index int) {
	i, j := 0, len(arr)
	for i < j {
		mid := i + (j-i)/2
		if num < arr[mid] {
			j = mid - 1
			continue
		} else if num > arr[mid] {
			i = mid + 1
			continue
		} else {
			return mid
		}
	}
	fmt.Printf("没找到")
	return 0
}

func main() {
	arr := []int{18, 22, 29, 35, 42, 50, 69}
	mid := binary_find(arr, 29)
	fmt.Printf("找到了：%d", mid)
}
