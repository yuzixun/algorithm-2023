package main

var res [][]int
var cur []int
var mark []bool

func permute(nums []int) [][]int {
	res, cur = [][]int{}, []int{}
	mark = make([]bool, len(nums))
	permuteHandle(nums)
	return res
}

func permuteHandle(nums []int) {
	if len(cur) == len(nums) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		res = append(res, tmp)
	} else {
		for i, m := range mark {
			if m {
				continue
			}

			cur = append(cur, nums[i])
			mark[i] = true
			permuteHandle(nums)
			mark[i] = false
			cur = cur[:len(cur)-1]
		}
	}
}
