package august

func maxSubArray(nums []int) int {
	maxValue := nums[0]

	for i := 1; i < len(nums); i++ {

		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}

		if maxValue < nums[i] {
			maxValue = nums[i]
		}
	}

	return maxValue
}
