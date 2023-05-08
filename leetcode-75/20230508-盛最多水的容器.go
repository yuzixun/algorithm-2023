package leetcode_75

func maxArea(height []int) int {
	i, j, res := 0, len(height)-1, 0

	for i < j {
		if height[i] < height[j] {
			res = max(res, height[i]*(j-i))
			i++
		} else {
			res = max(res, height[j]*(j-i))
			j--
		}
	}
	return res
}
