package main
// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。   
import (
	"fmt"
)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left<=right {
		mid := left +(right-left) / 2
		if target > nums[mid] {
			left = mid+1
		}else if target < nums[mid]{
			right = mid-1
		}else{
			fmt.Println("找到了")
			return mid
		}
	}
	fmt.Println("没找到")
	return -1
}

func main() {
	nums := []int{-1,0,3,5,9,12}
	target := 12
	fmt.Println(search(nums, target))
}
