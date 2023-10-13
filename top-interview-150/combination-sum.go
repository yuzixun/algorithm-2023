package main

var res [][]int
var cur []int
var used []bool

func combinationSum(candidates []int, target int) [][]int {
	cur, res = []int{}, [][]int{}
	used = make([]bool, len(candidates))
	combinationSumHandle(candidates, target, 0)
	return res
}

func combinationSumHandle(candidates []int, target, index int) {
	if target < 0 {
		return
	}
	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		res = append(res, tmp)
	} else {
		for i := index; i < len(candidates); i++ {
			cur = append(cur, candidates[i])
			combinationSumHandle(candidates, target-candidates[i], i)
			cur = cur[:len(cur)-1]
		}
	}
}
