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
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	reverse_tree(root.Left)
	reverse_tree(root.Right)
	return root
}

func levelorder(root *TreeNode) (res [][]int) {
	depth := 0
	var order func(root *TreeNode,depth int)
	order = func(root *TreeNode,depth int) {
		if root == nil {
			return 
		}
		if len(res)	== depth{
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], root.Val)
		order(root.Left,depth+1)
		order(root.Right,depth+1)
	}
	order(root,depth)
	// reverse(res)
    return res
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	node := create_tree(nums, 0)
	node = reverse_tree(node)
	res := levelorder(node)
	fmt.Print(res)
}
