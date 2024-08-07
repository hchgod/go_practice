package main

import (
	"fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0{
		return 0
	}
	res := make([][]int, len(obstacleGrid))
	for i, _ := range res {
		res[i] = make([]int, len(obstacleGrid[i]))
	}
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1{
			break
		}else{
			res[i][0] = 1
		}
	}
	for j := 0; j < len(obstacleGrid[0]); j++ {
		if obstacleGrid[0][j] == 1{
			break
		}else{
			res[0][j] = 1
		}
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[i]); j++{
			if obstacleGrid[i][j] == 1{
				res[i][j] = 0
			}else{
				res[i][j] = res[i-1][j] + res[i][j-1]
			}
		}
	}
	return res[len(res)-1][len(res[0])-1]
}

func main() {
	a := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(a))
}