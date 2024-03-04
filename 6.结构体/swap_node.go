package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func create_node(head *ListNode, val int) *ListNode {
	new_node := &ListNode{
		Val:  val,
		Next: nil,
	}
	head.Next = new_node
	return new_node
}

func create_list() *ListNode {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	cur := head
	for i := 2; i < 10; i++ {
		create_node(cur, i)
		cur = cur.Next
	}
	return head
}

func swapPairs(head *ListNode) *ListNode {
	temphead := &ListNode{}
	temphead.Next = head
	cur1 := temphead
	cur2 := cur1.Next
	for cur2 != nil&&cur2.Next != nil{
		cur1.Next = cur2.Next
		cur2.Next = cur2.Next.Next
		cur1.Next.Next = cur2

		cur1 = cur2
		cur2 = cur2.Next 
	}
	return temphead.Next
}


func main() {
	head := create_list()
	temphead := swapPairs(head)
	for ; temphead != nil; temphead = temphead.Next {
		fmt.Println(temphead.Val)
	}
}








