package main

import (
	"fmt"
)

var nums []int

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func CreateTreeNode(value int) *TreeNode {
	return &TreeNode{
		value: value,
		left:  nil,
		right: nil,
	}
}

/* 前序遍历 */
func preOrder(node *TreeNode) {

	if node == nil {
		return
	}
	// 访问优先级：根节点 -> 左子树 -> 右子树
	nums = append(nums, node.value)
	preOrder(node.left)
	preOrder(node.right)
}

/* 中序遍历 */
func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	// 访问优先级：左子树 -> 根节点 -> 右子树
	inOrder(node.left)
	nums = append(nums, node.value)
	inOrder(node.right)
}

/* 后序遍历 */
func postOrder(node *TreeNode) {
	if node == nil {
		return
	}
	nums = []int{}
	// 访问优先级：左子树 -> 右子树 -> 根节点
	postOrder(node.left)
	postOrder(node.right)
	nums = append(nums, node.value)
}

func main() {
	n1 := CreateTreeNode(1)
	n2 := CreateTreeNode(2)
	n3 := CreateTreeNode(3)
	n4 := CreateTreeNode(4)
	n5 := CreateTreeNode(5)
	// 构建节点之间的引用（指针）
	n1.left = n2
	n1.right = n3
	n2.left = n4
	n2.right = n5
	preOrder(n1)
	fmt.Println(nums)
}
