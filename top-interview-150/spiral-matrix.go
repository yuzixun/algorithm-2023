package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	var (
		minX, maxX, minY, maxY = 0, len(matrix[0]) - 1, 0, len(matrix) - 1
		result                 = make([]int, 0, len(matrix)*len(matrix[0]))
	)

	for {

		for x := minX; x <= maxX; x++ {
			result = append(result, matrix[minY][x])
		}
		minY++
		if minY > maxY {
			break
		}

		for y := minY; y <= maxY; y++ {
			result = append(result, matrix[y][maxX])
		}
		maxX--
		if minX > maxX {
			break
		}

		for x := maxX; x >= minX; x-- {
			result = append(result, matrix[maxY][x])
		}
		maxY--
		if minY > maxY {
			break
		}

		for y := maxY; y >= minY; y-- {
			result = append(result, matrix[y][minX])
		}

		minX++
		if minX > maxX {
			break
		}

	}

	return result
}
