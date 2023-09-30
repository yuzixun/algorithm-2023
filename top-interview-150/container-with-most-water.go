package main

import "github.com/algorithm-2023/tools"

func maxArea(height []int) int {
	left, right := 0, len(height)-1

	result := 0

	for left < right {
		if height[left] < height[right] {
			result = tools.Max(result, (right-left)*height[left])
			left++
		} else {
			result = tools.Min(result, (right-left)*height[right])
			right--
		}

	}
	return result
}
