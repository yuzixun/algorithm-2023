package leetcode_75

func maxOperations(nums []int, k int) int {
	res := 0
	m := map[int]int{}

	for _, num := range nums {
		if m[k-num] > 0 {
			m[k-num]--
			res++
		} else {
			m[num]++
		}
	}
	return res
}
