package main

func maxSubArray(nums []int) int {
	res := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[9]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}

	for i := 0; i < len(nums); i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
