package main

func searchMatrix(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix); i++ {
		left, right := 0, len(matrix[i])-1
		if target < matrix[i][left] || matrix[i][right] < target {
			continue
		}

		for left < right {
			mid := left + (right-left)/2
			switch {
			case matrix[i][mid] == target:
				return true
			case matrix[i][mid] < target:
				left = target + 1
			case matrix[i][mid] > target:
				right = target
			}
		}

	}
	return false
}
