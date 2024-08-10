package main
// 移除数组中重复的元素  删除有序数组中的重复项
import(
	"fmt"
)

func removeDuplicate(nums []int) int {
	slow := 0
	count := 1
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			count++
			slow++
			nums[slow] = nums[fast]
		}
	}
	nums = nums[:count]
	fmt.Println(nums)
	return len(nums[0:count])
}

func main() {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	nums_len := removeDuplicate(nums)
	fmt.Println(nums_len)
}