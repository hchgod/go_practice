package main

import "fmt"

func ModifySlice(s []int) {
    s[0] = 100 // 修改传入的 slice
    s = append(s, 200) // 添加元素，但这个改变不会影响原始 slice
	fmt.Println(s)
}

//slice：虽然传递的是对底层数组的引用，可以修改其内容，但对切片的重新分配（如使用 append）不会影响原始切片，

func main() {
    mySlice := []int{1, 2, 3}
    ModifySlice(mySlice)
    fmt.Println(mySlice) // 输出: [100 2 3]
}
