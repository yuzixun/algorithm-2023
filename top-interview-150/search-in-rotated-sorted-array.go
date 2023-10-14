package main

import "math"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// 落在左边
		if nums[0] <= target {
			if nums[mid] < nums[0] {
				nums[mid] = math.MaxInt64
			}
		} else {
			//	右边
			if nums[mid] >= nums[0] {
				nums[mid] = math.MinInt64
			}
		}

		if nums[mid] > nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
