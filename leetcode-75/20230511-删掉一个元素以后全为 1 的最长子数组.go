package leetcode_75

func longestSubarray(nums []int) int {
	left, right := 0, 0
	zeroCount := 0
	maxValue := 0

	for ; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		maxValue = max(maxValue, right-left)
	}

	return maxValue
}
