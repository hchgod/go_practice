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

func defs(left *TreeNode,right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}else if left==nil || right==nil{
		return false
	}else if left.Val != right.Val {
		return false
	}
	return defs(left.Left, right.Right) && defs(left.Right,right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return defs(root.Left,root.Right)
}

func isSymmetric1(root *TreeNode) bool {
	//用队列实现
    var queue []*TreeNode;
    if root != nil {
        queue = append(queue, root.Left, root.Right);
    }
    for len(queue) > 0 {
        left := queue[0];
        right := queue[1];
        queue = queue[2:];
        if left == nil && right == nil {
            continue;
        }
        if left == nil || right == nil || left.Val != right.Val {
            return false;
        };
        queue = append(queue, left.Left, right.Right, right.Left, left.Right);
    }
    return true;
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
	nums := []int{1, 2, 2, 3, 4, 4, 3}
	node := create_tree(nums, 0)
	res := isSymmetric(node)
	fmt.Print(res)
}
