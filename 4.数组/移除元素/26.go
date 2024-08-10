// 获取数组中去重后的长度   删除有序数组中的重复项
package main

import(
	"fmt"
)

func removeDuplicates(nums []int) int {
	slow := 0
	count := 1
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow = fast
			count++
			continue
		}
	}
	return count
}

func main() {
	nums := []int{1,1,2}
	count := removeDuplicates(nums)
	fmt.Println(count)
}