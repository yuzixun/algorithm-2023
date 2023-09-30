package main

func setZeroes(matrix [][]int) {
	is0X := false
	is0Y := false
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			is0X = true
			break
		}
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			is0Y = true
			break
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] != 0 {
			continue
		}
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = 0
		}
	}

	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] != 0 {
			continue
		}

		for i := 0; i < len(matrix); i++ {
			matrix[i][j] = 0
		}
	}

	if is0X {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	if is0Y {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
