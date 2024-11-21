package main

import "fmt"


// 二维数组
// func isSubsequence(s string, t string) bool {
// 	len1 := len(s)
// 	len2 := len(t)
// 	dp := make([][]bool, len1+1)
// 	for i := 0; i < len1+1; i++ {
// 		dp[i] = make([]bool, len2+1)
// 	}
// 	for i := 0; i < len2+1; i++ {
// 		dp[0][i] = true
// 	}

// 	for i := 1; i < len1+1; i++ {
// 		for j := 1; j < len2+1; j++ {
// 			if s[i-1] == t[j-1] {
// 				dp[i][j] = dp[i-1][j-1]
// 			} else {
// 				dp[i][j] = dp[i][j-1]
// 			}
// 		}
// 		fmt.Println(dp[i])
// 	}
// 	return dp[len1][len2]
// }

//一维数组
func isSubsequence(s string, t string) bool {
	len1 := len(s)
	len2 := len(t)
	dp := make([][]int, len2+1)


	for i := 1; i < len1+1; i++ {
		for j := 1; j < len2+1; j++ {
			if s[i-1] == t[j-1] {
				dp[j] = dp[i-1] + 1
			}
		}
		fmt.Println(dp)
	}
	return bool(dp[len2])
}


func main() {
	isSubsequence("axc", "ahbgdc")
}
