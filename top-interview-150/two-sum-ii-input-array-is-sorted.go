package main

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		s := numbers[left] + numbers[right]

		if s < target {
			left++
		} else if s > target {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return []int{-1, -1}
}
