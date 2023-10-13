package main

func exist(board [][]byte, word string) bool {
	var mark [][]bool
	for i := 0; i < len(board); i++ {
		mark = append(mark, make([]bool, len(board[i])))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if existHandle(board, word, mark, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func existHandle(board [][]byte, word string, mark [][]bool, index, i, j int) bool {
	if len(word) == index {
		return true
	}

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}

	if board[i][j] != word[index] {
		return false
	}
	if mark[i][j] {
		return false
	}

	mark[i][j] = true
	if existHandle(board, word, mark, index+1, i+1, j) {
		return true
	}
	if existHandle(board, word, mark, index+1, i-1, j) {
		return true
	}
	if existHandle(board, word, mark, index+1, i, j+1) {
		return true
	}
	if existHandle(board, word, mark, index+1, i, j-1) {
		return true
	}
	mark[i][j] = false

	return false
}
