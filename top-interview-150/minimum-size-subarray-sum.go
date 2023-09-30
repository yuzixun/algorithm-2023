package main

import "github.com/algorithm-2023/tools"

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minSize := 1 << 31
	left, right := 0, 0
	sum := 0

	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			minSize = tools.Min(minSize, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if minSize == 1<<31 {
		return 0
	}
	return minSize
}
