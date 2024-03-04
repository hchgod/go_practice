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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	    curA, curB := headA, headB
	    if curA == curB {
	        return curA
	    }
	    for curA != curB {
	        if curA == nil {
	            curA = headB
	        } else {
	            curA = curA.Next
	        }
	        if curB == nil {
	            curB = headA
	        } else {
	            curB = curB.Next
	        }
	        if curA == curB {
	            return curA
	        }
	    }
	    return nil
}

func main() {
	headA := create_list()
	headB := create_list()
	// value := 2
	temphead := getIntersectionNode(headA, headB)
	for ; temphead != nil; temphead = temphead.Next {
		fmt.Println(temphead.Val)
	}
}
