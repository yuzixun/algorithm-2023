package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index1, num := range nums {
		index2, exist := m[target-num]
		if exist {
			return []int{index1, index2}
		}
		m[num] = index1
	}
	return []int{-1, -1}
}
