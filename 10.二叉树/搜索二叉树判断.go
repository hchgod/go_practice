package main

// 使用队列层次遍历
import (
	"fmt"
	//"math"
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

func isValidBST(root *TreeNode) bool {
	//递归法实现
	if root == nil {
		return true
	}
	return check(root,root.Left.Val,root.Right.Val)
}

func check(node *TreeNode,min,max int) bool {
	if node == nil {
		return true
	}
	if min > node.Val || node.Val > max {
		return false
	}
	return check(node.Left,min,node.Val) && check(node.Right,node.Val,max)
}


func main() {
	nums := []int{2,1,3}
	node := create_tree(nums, 0)
	res := isValidBST(node)
	fmt.Print(res)
}
