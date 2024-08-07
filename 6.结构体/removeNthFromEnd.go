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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{}
	temp.Next = head
	slow := temp
	fast := slow.Next
	count := 1
	for fast != nil {
		if fast.Next == nil {
			slow.Next = slow.Next.Next
			return temp.Next
		}
		fast = fast.Next
		count++
		if count > n {
			slow = slow.Next
		}
	}
	return temp.Next
}

func main() {
	head := create_list()
	value := 2
	temphead := removeNthFromEnd(head, value)
	for ; temphead != nil; temphead = temphead.Next {
		fmt.Println(temphead.Val)
	}
}
