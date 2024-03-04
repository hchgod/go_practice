package main

import(
	"fmt"
	"reflect"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func test(){
	var list1 *ListNode
	fmt.Println(reflect.TypeOf(list1))
	list2 := &ListNode{
		Val: 1,
		Next: nil,
	}
	fmt.Print(reflect.TypeOf(list2))
}

func main() {
	test()
}