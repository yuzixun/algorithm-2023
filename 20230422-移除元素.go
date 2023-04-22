package main

func removeElement(nums []int, val int) int {
	var (
		slow, fast = 0, 0
	)

	for ; fast < len(nums); fast++ {
		if nums[fast] == val {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}

	return slow
}
