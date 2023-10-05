package main

import "sort"

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if isOverlap(cur, intervals[i]) {
			cur[0] = Min(cur[0], intervals[i][0])
			cur[1] = Max(cur[1], intervals[i][1])
		} else {
			res = append(res, cur)
			cur = intervals[i]
		}
	}
	res = append(res, cur)
	return res
}

func isOverlap(a, b []int) bool {
	return (a[0] <= b[0] && b[0] <= a[1]) || (a[0] <= b[1] && b[1] <= a[1]) || (b[0] <= a[0] && a[1] <= b[1])
}
