package alg

import "time"

const LIMIT = 8

func isValid(board *[8][8]int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == 1 {
			return false
		}
	}

	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	for i, j := row, col; i >= 0 && j < 8; i, j = i-1, j+1 {
		if board[i][j] == 1 {
			return false
		}
	}

	return true
}

func LDFS(board *[8][8]int, row int, startPos [2]int, delay time.Duration) bool {
	if row == 8 {
		return true
	}

	if row >= LIMIT {
		return false
	}

	colStart := 0
	if row == startPos[0] {
		colStart = startPos[1]
	}

	for col := colStart; col < 8; col++ {
		if isValid(board, row, col) {
			board[row][col] = 1

			time.Sleep(delay)
			if LDFS(board, row+1, startPos, delay) {
				return true
			}

			board[row][col] = 0
		}
	}

	return false
}
