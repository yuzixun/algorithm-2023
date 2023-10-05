package main

import "github.com/yuzixun/algorithm-2023/tools"

func longestConsecutive(nums []int) int {
	res := 0
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	for num := range set {
		if set[num-1] {
			continue
		}
		curCount := 0
		curNum := num
		for set[curNum] {
			curCount++
			curNum++
		}
		res = tools.Max(res, curCount)
	}

	return res
}
