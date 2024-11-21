package main

import "fmt"

//import "fmt"

func numDistinct(s string, t string) int {
	len1 := len(t)
	len2 := len(s)
	dp := make([][]int, len1+1)
	for i := 0; i < len1+1; i++ {
		dp[i] = make([]int, len2+1)
	}

	for j := 1; j < len2+1; j++ {
		if t[0] == s[j-1] {
			dp[1][j] = dp[1][j-1] + 1
		} else {
			dp[1][j] = dp[1][j-1]
		}
	}
	for i := 2; i < len1+1; i++ {
		for j := 1; j < len2+1; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			}else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	for i := 0; i < len1+1; i++ {
		fmt.Println(dp[i])
	}

	return dp[len1][len2]
}

// func numDistinct1(s string, t string) int {
//     dp:= make([][]int,len(s)+1)
//     for i:=0;i<len(dp);i++{
//         dp[i] = make([]int,len(t)+1)
//     }
//     // 初始化
//     for i:=0;i<len(dp);i++{
//         dp[i][0] = 1
//     }
//     // dp[0][j] 为 0，默认值，因此不需要初始化
//     for i:=1;i<len(dp);i++{
//         for j:=1;j<len(dp[i]);j++{
//             if s[i-1] == t[j-1]{
//                 dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
//             }else{
//                 dp[i][j] = dp[i-1][j]
//             }
//         }
//     }
// 	fmt.Println(dp[len(dp)-1][len(dp[0])-1])
//     return dp[len(dp)-1][len(dp[0])-1]
// }

func main() {
	// numDistinct("rabbbit","rabbit")
	numDistinct("babgbag", "bag")
	// numDistinct("abcde","ace")
}
