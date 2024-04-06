package main

// 使用队列层次遍历
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

func reverse_tree(root *TreeNode) (node *TreeNode) {
	//递归版本
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	reverse_tree(root.Right)
	reverse_tree(root.Left)
	return root
}

func reverse_tree2(root *TreeNode) *TreeNode {
	//前序遍历   用栈实现
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node.Left, node.Right = node.Right, node.Left
			node = node.Left
		}
		if node == nil {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		node = node.Right
	}
	return root
}

func reverse_tree3(root *TreeNode) *TreeNode {
	//中序遍历   用栈实现    不可实现
	stack := []*TreeNode{}
	node := root
	for node != nil || stack != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		node.Left,node.Right = node.Right,node.Left
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return root
}

func levelorder(root *TreeNode) (res [][]int) {
	depth := 0
	var order func(root *TreeNode, depth int)
	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(res) == depth {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], root.Val)
		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}
	order(root, depth)
	// reverse(res)
	return res
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	node := create_tree(nums, 0)
	node = reverse_tree3(node)
	res := levelorder(node)
	fmt.Print(res)
}
