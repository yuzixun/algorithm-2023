package leetcode_75

func findMaxAverage(nums []int, k int) float64 {
	sumValue := sum(nums[:k])
	maxValue := sumValue

	for i := k; i < len(nums); i++ {
		sumValue = sumValue - nums[i-k] + nums[i]
		maxValue = max(maxValue, sumValue)
	}

	return float64(maxValue) / float64(k)
}
