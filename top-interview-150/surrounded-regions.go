package main

func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			isEdge := i == 0 || j == 0 || i == len(board)-1 || j == len(board[0])-1
			if isEdge && board[i][j] == 'O' {
				solveHandle(board, i, j)
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == '#' {
				board[i][j] = 'O'
			}

		}
	}
}

func solveHandle(board [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return
	}
	if board[i][j] == 'X' || board[i][j] == '#' {
		return
	}

	board[i][j] = '#'
	solveHandle(board, i-1, j)
	solveHandle(board, i, j-1)
	solveHandle(board, i+1, j)
	solveHandle(board, i, j+1)
}
