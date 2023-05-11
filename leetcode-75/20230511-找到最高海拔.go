package leetcode_75

func largestAltitude(gain []int) int {
	var (
		maxValue = 0
		value    = 0
	)

	for _, item := range gain {
		value += item
		maxValue = max(maxValue, value)
	}

	return maxValue
}
