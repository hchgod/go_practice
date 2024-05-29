package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return len(nums)
	}else if n == 2{
		if nums[1] - nums[0] !=0{
			return 2
		} else {
			return 1
		}
	}
	res := 2
	for i, j := 0, 1; j < len(nums)-1; {
		prediff := nums[i+1] - nums[i]
		postdiff := nums[j+1] - nums[j]
		//fmt.Println(prediff, postdiff)
		if (prediff * postdiff) <= 0 {
			res++
			j++
			i++
			continue
		}
		j++
	}
	fmt.Println(res)
	return res
}

func main() {
	list := []int{1,17,5,10,13,15,10,5,16,8}
	wiggleMaxLength(list)
}
