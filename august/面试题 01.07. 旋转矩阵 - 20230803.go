package august

func rotate(matrix [][]int) {

	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i] = matrix[n-i], matrix[i]
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
