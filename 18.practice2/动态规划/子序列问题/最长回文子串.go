package main
// 647. 回文子串
import "fmt"

func longestPalindromeSubseq(s string) string {
	res := 0
	maxl := 1
	maxstr := s[:1]
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if (i == j) {
				dp[i][j] = true
				res++
			}else if s[i] == s[j] {
				if (i > j + 1){
					dp[i][j] = dp[i-1][j+1]
				} else {
					dp[i][j] = true
				}
				if dp[i][j] == true {
					if maxl < i-j+1 {
						maxl = i-j+1
						maxstr = s[j:i+1]
					}
					fmt.Println(maxstr)
					res++
				}
			}
		}
		fmt.Println(dp[i])
	}
	return maxstr
}

func main() {
	fmt.Println(longestPalindromeSubseq("eu"))
}
