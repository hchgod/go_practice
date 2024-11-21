package main
// 647. 回文子串
import "fmt"

func countSubstrings(s string) int {
	res := 0
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
					res++
				}
			}
		}
		fmt.Println(dp[i])
	}
	return res
}

func main() {
	fmt.Println(countSubstrings("abba"))
}
