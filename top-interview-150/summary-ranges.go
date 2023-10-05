package main

import "fmt"

func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return res
	}

	start := nums[0]
	end := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != end+1 {
			res = append(res, formatString(start, end))
			start, end = nums[i], nums[i]
		} else {
			end = nums[i]
		}
	}
	res = append(res, formatString(start, end))
	return res
}

func formatString(start, end int) string {
	if start == end {
		return fmt.Sprint(start)
	}
	return fmt.Sprintf("%d->%d", start, end)
}
