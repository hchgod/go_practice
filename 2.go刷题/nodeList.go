package main

import (
	"fmt"
)

type ListNode struct {
	value int
	next  *ListNode
}

func NewListNode(num int) *ListNode {
	return &ListNode{
		value: num,
		next:  nil,
	}
}

func insertNode(node1 *ListNode, P *ListNode) {
	next := node1.next
	node1.next = P
	P.next = next
}

// 打印节点的值
func print_node(node *ListNode) {
	for i := node; i != nil; {
		fmt.Printf("%d\n", i.value)
		i = i.next
	}
}

// 删除节点之后的首个节点
func delete_node(node *ListNode) {
	if node.next == nil {
		return
	}
	next := node.next.next
	node.next.next = nil
	node.next = next
}

func main() {
	node1 := NewListNode(32)
	node2 := NewListNode(33)
	node3 := NewListNode(34)
	node4 := NewListNode(35)
	node5 := NewListNode(36)
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	insertNode(node1, NewListNode(39))
	delete_node(node3)
	print(node1)
}
