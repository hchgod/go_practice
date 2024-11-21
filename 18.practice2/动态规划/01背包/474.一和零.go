package main
//三维数组
import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(strs); i++ {
		m_0 := m
		for ;m_0 >= 0 ;m_0--{
			n_0 := n
			for ;n_0 >= 0;n_0-- {
				count_0, count_1 := count(strs[i])
				if count_0 <= m_0 && count_1 <= n_0 {
					dp[m_0][n_0] = max(dp[m_0][n_0], dp[m_0-count_0][n_0-count_1]+1)
				}
			}
		}
		fmt.Println(dp[:][:])
	}
	return dp[m][n]
}

func main() {
	findMaxForm([]string{"10","0","1"},1,1)
}

func count(str string) (m0 int, n1 int) {
	for _, value := range str {
		if string(value) == "0" {
			m0++
		} else {
			n1++
		}
	}
	return m0, n1
}
