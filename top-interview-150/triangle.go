package main

func minimumTotal(triangle [][]int) int {
	var (
		y = len(triangle)
	)
	dp := make([][]int, y)
	for i := 0; i < y; i++ {
		dp[i] = make([]int, len(triangle[i]))
		if i == y-1 {
			dp[i] = triangle[i]
		}
	}

	for i := y - 2; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			dp[i][j] = Min(dp[i+1][j+1], dp[i+1][j]) + triangle[i][j]
		}
	}

	return dp[0][0]
}
