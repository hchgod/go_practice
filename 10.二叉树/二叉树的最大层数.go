package main

// 使用队列层次遍历
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

func max_level(root *TreeNode) (result int) {
	//是用二维数组表示
	res := [][]int{}
	if root == nil {
		return result
	}
	queue := list.New()
	queue.PushBack(root)
	depth := 0
	for ; queue.Len() > 0; depth++ {
		res = append(res, []int{})
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node == nil {
				continue
			}
			res[depth] = append(res[depth], node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	result = len(res)
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	node := create_tree(nums, 0)
	res := max_level(node)
	fmt.Print(res)
}
