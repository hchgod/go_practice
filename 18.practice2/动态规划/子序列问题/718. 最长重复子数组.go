package main

//返回两个数组中公共的、长度最长的子数组的长度  连续的。 
import "fmt"

//动态规划的解法
func findLength(nums1 []int, nums2 []int) int {
	maxvalue := 0
	len1 := len(nums1)
	len2 := len(nums2)
	dp := make([][]int, len1+1)
	for i := 0; i < len1+1; i++ {
		dp[i] = make([]int, len2+1)
	}

	for i := 1; i < len1+1; i++ {
		for j := 1; j < len2+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > maxvalue {
				maxvalue = dp[i][j]
			}
		}
	}
	fmt.Println(dp)
	return maxvalue
}

func main() {
	findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7})
}
