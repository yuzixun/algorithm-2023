package main

func lengthOfLIS(nums []int) int {
	arr := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		left, right := 0, len(arr)-1
		for left <= right {
			mid := left + (right-left)/2
			switch {
			case arr[mid] == nums[i]:
				right = mid - 1
			case arr[mid] > nums[i]:
				right = mid - 1
			case arr[mid] < nums[i]:
				left = mid + 1
			}
		}

		if left >= len(arr) {
			arr = append(arr, nums[i])
		} else {
			arr[left] = nums[i]
		}
	}

	return len(arr)
}
