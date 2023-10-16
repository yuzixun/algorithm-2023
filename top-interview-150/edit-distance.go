package main

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 1; i <= len(word1); i++ {
		dp[0][i] = i
	}
	for i := 1; i <= len(word2); i++ {
		dp[i][0] = i
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(dp[i][j-1], Min(dp[i-1][j-1], dp[i-1][j]))
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
