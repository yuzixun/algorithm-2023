package main

func generateParenthesis(n int) []string {
	var res []string

	var dfs func(left, right int, cur string)
	dfs = func(left, right int, cur string) {
		if left == 0 && right == 0 {
			res = append(res, cur)
			return
		}
		if left > 0 {
			dfs(left-1, right, cur+"(")
		}
		if left < right {
			dfs(left, right-1, cur+")")
		}
	}

	dfs(n, n, "")
	return res
}
