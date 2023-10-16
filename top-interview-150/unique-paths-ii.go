package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([]int, len(obstacleGrid[0]))

	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j] == 0 {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[len(dp)-1]
}
