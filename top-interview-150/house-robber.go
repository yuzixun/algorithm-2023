package main

func rob(nums []int) int {
	pre := 0
	cur := nums[0]
	for i := 2; i <= len(nums); i++ {
		cur, pre = max(cur, pre+nums[i-1]), cur
	}
	return cur
}
