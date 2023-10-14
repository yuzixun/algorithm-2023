package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	l, r := -1, -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] == target:
			right = mid - 1
		case nums[mid] > target:
			right = mid - 1
		case nums[mid] < target:
			left = mid + 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	l = left

	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] == target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		case nums[mid] < target:
			left = mid + 1
		}
	}
	r = right

	return []int{l, r}
}
