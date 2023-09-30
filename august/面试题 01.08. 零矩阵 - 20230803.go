package august

func setZeroes(matrix [][]int) {
	x, y := len(matrix[0]), len(matrix)
	x0, y0 := false, false

	for i := 0; i < x; i++ {
		if matrix[0][i] == 0 {
			x0 = true
			break
		}
	}

	for i := 0; i < y; i++ {
		if matrix[i][0] == 0 {
			y0 = true
			break
		}
	}

	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			if matrix[j][i] == 0 {
				matrix[j][0] = 0
				matrix[0][i] = 0
			}
		}
	}

	for i := 1; i < x; i++ {
		if matrix[0][i] == 0 {
			for j := 1; j < y; j++ {
				matrix[j][i] = 0
			}
		}
	}

	for i := 1; i < y; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < x; j++ {
				matrix[i][j] = 0
			}
		}
	}

	if x0 {
		for i := 0; i < x; i++ {
			matrix[0][i] = 0
		}
	}

	if y0 {
		for i := 0; i < y; i++ {
			matrix[i][0] = 0
		}
	}
}
