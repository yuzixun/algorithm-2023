package leetcode_75

func sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}
