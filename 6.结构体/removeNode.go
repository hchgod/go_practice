package main

import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func create_node(head *ListNode,val int) *ListNode{
	new_node := &ListNode{
		Val: val,
		Next: nil,
	}
	head.Next = new_node
	return new_node
}

func create_list() *ListNode{
	head := &ListNode{
		Val: 1,
		Next: nil,
	}
	cur := head
	for i:=2; i<10; i++{
		create_node(cur,i)
		cur = cur.Next
	}
	return head
}

 func removeElements(head *ListNode, val int) *ListNode {
    temphead := &ListNode{}
	temphead.Next = head
	cur:=temphead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}else{
			cur = cur.Next
		}
	}
	return temphead.Next
}

func main() {
	head := create_list()
	temphead := removeElements(head,3)
	for ;temphead != nil;temphead=temphead.Next{
		fmt.Println(temphead.Val)	
	}
}