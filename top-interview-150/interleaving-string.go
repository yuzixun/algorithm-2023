package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1+l2 != len(s3) {
		return false
	}
	dp := make([][]bool, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]bool, l2+1)
	}
	for i := 0; i <= l1 && s1[i] == s3[i]; i++ {
		dp[i][0] = true
	}
	for j := 0; j <= l2 && s2[j] == s3[j]; j++ {
		dp[0][j] = true
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if (s3[i+j-1] == s2[j] && dp[i][j-1]) || s3[i+j-1] == s1[i] && dp[i-1][j] {
				dp[i][j] = true
			}
		}
	}

	return dp[l1][l2]
}
