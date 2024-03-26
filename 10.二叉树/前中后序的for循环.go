package main

import (
	"container/list"
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
	if root == nil {
		return res
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode) //断言可以强制类型转换
		res = append(res, node.Val)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return res
}

func reverse(res []int) {
	i,j := 0, len(res)-1
	for i < j{
		res[i],res[j] = res[j],res[i]
		i++
		j--
	}
}

func postorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return res
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, node.Val)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}
	reverse(res)
	return res
}

func midorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return res
	}
	stack := list.New()
	stack.PushBack(root)
	cur := root
	for cur != nil||stack.Len()>0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}else{
			node := stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, node.Val)
			cur = node.Right
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	node := create_tree(nums, 0)
	res := postorderTraversal(node)
	fmt.Print(res)
}
