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

func levelOrder1(root *TreeNode) (res []int) {
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	// depth := 0
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		if node == nil {
			continue
		}
		res= append(res,node.Val)
		queue.PushBack(node.Left)
		queue.PushBack(node.Right)
	}
    return res
}


func levelOrder2(root *TreeNode) (res [][]int) {
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	var tmparr []int
	for queue.Len() > 0 {
		res = append(res, []int{}}
		for i:=0; i<len(res[]); i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node == nil {
				continue
			}
			res= append(res,node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmparr
		}
		
	}
    return res
}

func reverse(res [][]int){
	i,j := 0,len(res)-1
	for i < j{
		res[i],res[j] = res[j],res[i]
		i++
		j--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	node := create_tree(nums, 0)
	res := levelOrder1(node)
	fmt.Print(res)
	// node.preorder()
}
