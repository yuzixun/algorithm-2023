package leetcode_75

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	first, second := nums[0], 1<<31

	for i := 1; i < len(nums); i++ {
		if nums[i] > second {
			return true
		} else if nums[i] > first {
			second = nums[i]
		} else {
			first = nums[i]
		}
	}

	return false
}
