package main

import "fmt"

func function() int {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	dp := make([][]int, len(obstacleGrid))
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(obstacleGrid[i]))
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[0][i] == 1{
			break
		}
		dp[0][i] = 1
	}
	for i := 0; i < n; i++ {
		if obstacleGrid[i][0] == 1{
			break
		}
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if obstacleGrid[i][j] == 1{
				dp[i][j] = 0
			}else{
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	fmt.Println(dp)
	return dp[n-1][m-1]
}

func main() {
	function()
}
