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

func detectCycle(head *ListNode) *ListNode {
	slow,fast := head,head
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for head != slow{
				slow = slow.Next
				head = head.Next
			}
			return head
		}		
	}
	return nil
}


func main() {
	headA := create_list()
	// value := 2
	temphead := detectCycle(headA)
	for ; temphead != nil; temphead = temphead.Next {
		fmt.Println(temphead.Val)
	}
}
