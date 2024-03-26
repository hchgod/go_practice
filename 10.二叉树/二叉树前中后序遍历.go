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

func preorderTraversal(root *TreeNode) (res []int) {
	var Traversal func(root *TreeNode)
    Traversal = func(root *TreeNode){
        if root== nil{
            return
        }
        res = append(res, root.Val)
        Traversal(root.Left)
        Traversal(root.Right)
    }
    Traversal(root)
    return res
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	node := create_tree(nums, 0)
	res := preorderTraversal(node)
	fmt.Print(res)
	// node.preorder()
}
