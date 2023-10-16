package main

import "math"

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix[0]); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	res := math.MinInt
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if i == 0 || j == 0 {
				dp[i][j] = 1
			}

			if dp[i-1][j] == '1' && dp[i][j-1] == '1' {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 1
			}
			res = max(dp[i][j], res)
		}
	}
	return res
}
