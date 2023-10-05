package main

func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]bool)

	for i, num := range nums {
		if i-k-1 >= 0 {
			set[nums[i-k-1]] = false
		}
		if set[num] {
			return true
		}
		set[num] = true
	}
	return false
}
