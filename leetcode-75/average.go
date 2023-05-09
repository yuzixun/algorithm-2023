package leetcode_75

func average(nums []int, k int) float64 {
	var sum int
	for _, num := range nums {
		sum += num
	}

	return float64(sum) / float64(k)
}
