package main

import (
	"fmt"
)

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func TestListNode(){
	head := new(ListNode)
	head.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 3
	ln4 := new(ListNode)
	ln4.Val = 4
	ln5 := new(ListNode)
	ln5.Val = 5
	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next=ln5
	fmt.Println(reverseList(head))
	// Println(reverseListV2(head))
}
func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}


func main() {
	TestListNode()
}
