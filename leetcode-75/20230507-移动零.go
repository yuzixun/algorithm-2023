package leetcode_75

func moveZeroes(nums []int) {
	i := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[i] = nums[fast]
			i++
		}
	}

	for ; i < len(nums); i++ {
		nums[i] = 0
	}
}
