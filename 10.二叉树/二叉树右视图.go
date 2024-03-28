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

func levelOrder2(root *TreeNode) (res [][]int) {
	//是用二维数组表示
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	// var tmparr []int
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
	res = right(res)
	return res
}


func right(res [][]int) [][]int { //右视图
	result := [][]int{}
	for i := 0; i < len(res); i++ {
		result[0] = append(result[0], res[i][len(res[i])-1])
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	node := create_tree(nums, 0)
	res := levelOrder2(node)
	fmt.Print(res)
}
