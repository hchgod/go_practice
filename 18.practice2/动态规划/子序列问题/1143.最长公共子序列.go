package main

//返回两个数组中公共的、长度最长的子数组的长度。 dp数组表示以i为结尾的递增子数组长度
import "fmt"

//动态规划的解法
func longestCommonSubsequence(text1 string, text2 string) int {
	maxvalue := 0
	len1 := len(text1)
	len2 := len(text2)
	dp := make([][]int, len1+1)
	for i := 0; i < len1+1; i++ {
		dp[i] = make([]int, len2+1)
	}

	for i := 1; i < len1+1; i++ {
		for j := 1; j < len2+1; j++ {
			if string(text1[i-1]) == string(text2[j-1]) {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			if dp[i][j] > maxvalue {
				maxvalue = dp[i][j]
			}
		}
	}
	for i := 0; i < len1+1; i++ {
		fmt.Println(dp[i])
	}

	return maxvalue
}

func main() {
	longestCommonSubsequence("bsbininm","jmjkbkjkv")
	longestCommonSubsequence("abcde","ace")
}
