package main

func isValidSudoku(board [][]byte) bool {
	var (
		x = make([]int, 9, 9)
		y = make([]int, 9, 9)
		s = make([]int, 9, 9)
	)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}

			if (x[i]>>(board[i][j]-'0'))&1 == 1 {
				return false
			}
			x[i] = x[i] | (1 << (board[i][j] - '0'))

			if (y[j]>>(board[i][j]-'0'))&1 == 1 {
				return false
			}
			y[j] = y[j] | (1 << (board[i][j] - '0'))

			xt, yt := i/3, j/3
			index := xt*3 + yt

			if (s[index]>>(board[i][j]-'0'))&1 == 1 {
				return false
			}
			s[index] = s[index] | (1 << (board[i][j] - '0'))
		}
	}

	return true
}
