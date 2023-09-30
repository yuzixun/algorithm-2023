package main

func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j]%10 == 1 {
				handleXY(board, j, i)
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			liveCount := board[i][j] / 10
			liveStatus := board[i][j] % 10
			if liveCount == 2 && liveStatus == 1 {
				board[i][j] = 1
			} else if liveCount == 3 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
}

func handleXY(board [][]int, x, y int) {
	for deltaX := -1; deltaX <= 1; deltaX++ {
		for deltaY := -1; deltaY <= 1; deltaY++ {
			cX, cY := x+deltaX, y+deltaY

			if cX < 0 || cX >= len(board[0]) || cY < 0 || cY >= len(board) || (cX == x && cY == y) {
				continue
			}
			board[cY][cX] = 10 + board[cY][cX]
		}
	}
}
