package main

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	if nums[left] <= nums[right] {
		return nums[left]
	}

	for left <= right {
		mid := left + (right-left)/2

		if nums[0] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return left
}
