package leetcode_75

func longestOnes(nums []int, k int) int {
	var (
		left, right int
		maxValue    int
	)

	for right < len(nums) {

		if nums[right] == 0 {
			if k == 0 {
				for nums[left] == 1 {
					left++
				}
				left++
			} else {
				k--
			}
		}

		maxValue = max(maxValue, right-left+1)
		right++
	}

	return maxValue
}
