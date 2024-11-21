package main
// 647. 回文子串
import "fmt"

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}

	for i := 0; i < len(s); i++ {
		for j := i; j >= 0; j-- {
			if s[i] == s[j] {
				if (i > j + 1){
					dp[i][j] = dp[i-1][j+1] + 2
				} else if i == j {
					dp[i][j] = 1
				} else if (i == j + 1){
					dp[i][j] = 2
				}
			} else {
				dp[i][j] = max(dp[i][j+1], dp[i-1][j])
			}
		}
		fmt.Println(dp[i])
	}
	return dp[len(s)-1][0]
}

func longestPalindromeSubseq1(s string) int {
	size := len(s)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
		dp[i][i] = 1
	}
	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][size-1]
}

func main() {
	fmt.Println(longestPalindromeSubseq1("euffse"))
	fmt.Println(longestPalindromeSubseq("exffse"))
}