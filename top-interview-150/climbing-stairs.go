package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	var (
		cur int
		v1  = 1
		v2  = 1
	)
	for i := 1; i <= n; i++ {
		cur, v1, v2 = v1+v2, cur, v1
	}
	return cur
}
