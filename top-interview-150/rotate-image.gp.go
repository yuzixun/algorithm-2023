package main

func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	size := len(matrix)
	for y := 0; y < size/2; y++ {
		for x := y; x < size-y-1; x++ {
			matrix[y][x], matrix[x][size-y-1], matrix[size-y-1][size-1-x], matrix[size-1-x][y] = matrix[size-1-x][y], matrix[y][x], matrix[x][size-y-1], matrix[size-y-1][size-1-x]
		}
	}

}
