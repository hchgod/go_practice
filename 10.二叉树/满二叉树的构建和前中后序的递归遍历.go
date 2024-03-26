package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func create(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func create_tree(nums []int, index int) *TreeNode {
	if index < len(nums) {
		node := create(nums[index])
		node.Left = create_tree(nums, 2*index+1)
		node.Right = create_tree(nums, 2*index+2)
		return node
	}
	return nil
}

func preorderTraversal(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = preorderTraversal(root.Left, res)
	res = preorderTraversal(root.Right, res)
	return res
}

func midorderTraversal(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = midorderTraversal(root.Left, res)
	res = append(res, root.Val)
	res = midorderTraversal(root.Right, res)
	return res
}

func postorderTraversal(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = postorderTraversal(root.Left, res)
	res = postorderTraversal(root.Right, res)
	res = append(res, root.Val)
	return res
}

// func (node *TreeNode) preorder() {
// 	if node == nil {
// 		return
// 	}
// 	fmt.Println(node.Val)
// 	node.Left.preorder()
// 	node.Right.preorder()
// }

// func (node *TreeNode) midorder() {
// 	if node == nil {
// 		return
// 	}
// 	node.Left.midorder()
// 	fmt.Println(node.Val)
// 	node.Right.midorder()
// }

// func (node *TreeNode) postorder() {
// 	if node == nil {
// 		return
// 	}
// 	node.Left.postorder()
// 	node.Right.postorder()
// 	fmt.Println(node.Val)
// }

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	node := create_tree(nums, 0)
	res := []int{}
	res = postorderTraversal(node, res)
	fmt.Print(res)
	// node.preorder()
}
